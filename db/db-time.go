package db

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"timetracker/helpers"
	"timetracker/models"
)

// Describes the structure of fields on a time entry that can be updated
type UpdatedTimeEntry struct {
	Comments  *string            `json:"comments"`
	Value     *float32           `json:"value"`
	ValueType *string            `json:"valueType"`
	Tags      *[]models.Tag      `json:"tags,omitempty"`
	RepoItems *[]models.RepoItem `json:"repoItems,omitempty"`
}

// Columns to pull when getting time data
var timecols = []string{
	"id",
	"created",
	"updated",
	"value",
	"valueType",
	"comments",
	"userId",
	"username",
	"userAvatar",
	"organisationId",
	"organisation",
	"organisationAvatar",
}

// Gets all time entries for a user and the given date range
// returning the available rows and a page of data for the given date range
func GetTimeEntries(userId uint, dateFrom, dateTo time.Time, page Pagination) (uint, []models.TimeEntry) {
	dbConn := ConnectDB()

	// Middle section of our SQL query, will be constructed when we make requests
	queryMid := ` FROM vw_time_entries AS t WHERE t.userId = $1 AND t.created >= $2 AND t.created <= $3`

	// initialise the return values
	rowCount := uint(0)
	time := []models.TimeEntry{}

	// construct the query for getting the available rows
	// then query the db and parse the count from the returned row
	query := "SELECT COUNT(*) " + queryMid
	row := dbConn.QueryRow(query, userId, dateFrom, dateTo)
	err := row.Scan(&rowCount)

	// Handle any errors returned from our COUNT query
	helpers.HandleError(err)

	if rowCount > 0 {
		// if the column to sort by is not in the global timecols array
		// default to sorting by the creation date & time of the time entries
		if !helpers.StrArrayContains(timecols, page.Sort) {
			page.Sort = "created"
		}

		// construct our sorting and pagination sections of the query
		queryEnd := fmt.Sprintf(" ORDER BY %s %s LIMIT %d OFFSET %d;", page.Sort, page.SortDirection(), page.GetPageSize(), page.Offset())

		// Construct the query to retreive a page of time entries
		query = "SELECT " + strings.Join(timecols, ", ") + queryMid + queryEnd

		// Get any avaiable time entries for the given query
		time = getTimeEntries(query, userId, dateFrom, dateTo)
	}

	// return parsed rows
	return rowCount, time
}

// Gets the time entry with a given id
func GetTimeEntry(id uint) models.TimeEntry {
	// Create our parameterised SQL query using our vw_time_entries view
	query := "SELECT " + strings.Join(timecols, ", ") + " FROM vw_time_entries AS t WHERE t.id = $1;"
	timeEntries := getTimeEntries(query, id)

	var timeEntry models.TimeEntry

	if len(timeEntries) > 0 {
		timeEntry = timeEntries[0]

		// Set the RepoItems and Tags properties to empty arrays
		timeEntry.Tags = getTimeEntryTags(id)
		timeEntry.RepoItems = getTimeEntryRepos(id)
	}

	// Return the parsed time entry
	return timeEntry
}

// Gets the time entries for the given prepared sql query and parameters
func getTimeEntries(query string, args ...interface{}) []models.TimeEntry {
	// Query the db with the above statement and handle any returned errors
	dbConn := ConnectDB()
	rows, err := dbConn.Query(query, args)
	helpers.HandleError(err)

	// defer closing the sql.Rows cursor until the function has finished
	defer rows.Close()

	// Array of our time entries that will be returned
	time := []models.TimeEntry{}

	// Loop through each data row returned and parse it into a time entry
	for rows.Next() {
		var timeEntry models.TimeEntry
		var user models.OwnerTrimmed
		var org models.OwnerTrimmed

		// Convert the columns returned in the row of data into a time entry
		rows.Scan(
			&timeEntry.Id,
			&timeEntry.Created,
			&timeEntry.Updated,
			&timeEntry.Value,
			&timeEntry.ValueType,
			&timeEntry.Comments,
			&user.Id,
			&user.Name,
			&user.Avatar,
			&org.Id,
			&org.Name,
			&org.Avatar)

		// Set the user and organisation values on the time entry to those parsed above
		timeEntry.User = user
		timeEntry.Organisation = org

		// add our parsed time entry into the retulting array
		time = append(time, timeEntry)
	}

	return time
}

// Gets the tags linked to the time entry
func getTimeEntryTags(timeEntryId uint) []models.Tag {
	tags := []models.Tag{}

	// Get the tags for the time entry and handle any returned errors
	dbConn := ConnectDB()
	query := `SELECT id, name FROM vw_time_entry_tags WHERE timeentryid = $1`
	tagRows, err := dbConn.Query(query, timeEntryId)
	helpers.HandleError(err)

	// Defer closing of the tagRows cursor till the function returns
	defer tagRows.Close()

	// loop through the rows returned and parse them to an instance of Tag
	// and append each value to timeEntry.Tags
	for tagRows.Next() {
		var tag models.Tag
		tagRows.Scan(&tag.Id, &tag.Name)
		tags = append(tags, tag)
	}

	return tags
}

// Gets the repo items linked to the selected time entry
func getTimeEntryRepos(timeEntryId uint) []models.RepoItem {
	repoItems := []models.RepoItem{}
	query := `
		SELECT id, created, updated, itemIdSource, itemType, source, repoName, description
		FROM vw_repo_items 
		WHERE timeentryid = $1;
	`
	// Get the repo items added to the time entry and handle any errors
	dbConn := ConnectDB()
	repoRows, err := dbConn.Query(query, timeEntryId)
	helpers.HandleError(err)
	defer repoRows.Close()

	// loop through the rows returned and parse them to an instance of RepoItem
	// and append each value to timeEntry.RepoItems
	for repoRows.Next() {
		var repoItem models.RepoItem
		repoRows.Scan(
			&repoItem.Id,
			&repoItem.Created,
			&repoItem.Updated,
			&repoItem.ItemIdSource,
			&repoItem.ItemType,
			&repoItem.Source,
			&repoItem.RepoName,
			&repoItem.Description)

		repoItems = append(repoItems, repoItem)
	}

	return repoItems
}

// Creates a new time entry in the database from the passed in models.TimeEntry
// and returns the id of the new time entry
func CreateTimeEntry(newEntry models.TimeEntry) uint {
	// Connect to the DB
	dbConn := ConnectDB()

	// sp_time_insert stored procedure returns the id of the new time entry
	// so call QueryRow with a prepared statement to just grab this value
	// QueryRow returns a pointer to an instance of sql.Row with a single value
	row := dbConn.QueryRow(`call sp_time_insert($1, $2, $3, $4, $5, 0)`,
		newEntry.User.Id,
		newEntry.Organisation.Id,
		newEntry.Comments,
		newEntry.Value,
		newEntry.ValueType)

	// Parse the new time entry id returned by the db in the row variable
	// and handle any returned errors
	err := row.Scan(&newEntry.Id)
	helpers.HandleError(err)

	// Insert any tags linked to the time entry
	insertTags(&newEntry.Tags, newEntry.Id, newEntry.User.Id)

	// Insert any repo items for the new time entry
	insertRepoItem(&newEntry.RepoItems, newEntry.Id)

	// Return the time entries new id
	return newEntry.Id
}

// Updates an existing time entry
func UpdateTimeEntry(userId, timeEntryId uint, vals UpdatedTimeEntry) error {
	if err := checkUserIdAndTimeEntryIdValue(userId, timeEntryId); err != nil {
		return err
	}

	// Connect to the database and defer closing until the func ends
	dbConn := ConnectDB()

	if vals.Comments != nil {
		updateTimeProp(timeEntryId, "comments", vals.Comments)
	}
	if vals.Value != nil {
		updateTimeProp(timeEntryId, "value", vals.Value)
	}
	if vals.ValueType != nil {
		updateTimeProp(timeEntryId, "valueType", vals.ValueType)
	}
	if vals.Tags != nil {
		_, err := dbConn.Exec(`DELETE FROM tbl_timeentrytaglinks WHERE timeentryid = $1`, vals.ValueType, timeEntryId)
		helpers.HandleError(err)
		insertTags(vals.Tags, timeEntryId, userId)
	}
	if vals.RepoItems != nil {
		rows, err := dbConn.Query(`SELECT id FROM tbl_repoitem WHERE timeEntryId = $1`, timeEntryId)
		helpers.HandleError(err)
		defer rows.Close()

		for rows.Next() {
			var id uint
			rows.Scan(&id)
			isInList := false

			// Check if the item is in the repo items list
			for _, r := range *vals.RepoItems {
				if r.Id == id {
					isInList = true
					break
				}
			}

			// Delete the item if its not in the list
			if !isInList {
				dbConn.Exec(`call sp_time_repoitem_delete($1)`, id)
			}
		}

		// Add in any new repo items
		insertRepoItem(vals.RepoItems, timeEntryId)
	}

	return nil
}

// Links the tags to the time entry, inserting any new user created tags into the tags table
func insertTags(tags *[]models.Tag, timeEntryId, userId uint) {
	// loop through the tags added to the entry and add the new tags
	// and/or the link between the tags and time entry using our
	// sp_time_tags_insert stored procedure
	dbConn := ConnectDB()
	query := `call sp_time_tags_insert($1, $2, $3, $4)`
	for _, tag := range *tags {
		dbConn.Exec(query, timeEntryId, tag.Name, tag.Id, userId)
	}
}

// Adds any new repo items to the selected time entry
func insertRepoItem(repoItems *[]models.RepoItem, timeEntryId uint) {
	dbConn := ConnectDB()
	query := `call sp_time_repoitem_insert($1, $2, $3, $4, $5, $6, $7)`

	// Loop through the linked repository items to the database with a link to this
	// new time entry, using the sp_time_repoitem_insert stored procedure
	for _, r := range *repoItems {
		if r.Id == 0 {
			dbConn.Exec(query, timeEntryId, r.Created, r.ItemIdSource, r.ItemType, r.Source, r.RepoName, r.Description)
		}
	}
}

// Updates the selected property on the time entry with the new value
// handles any errors returned from the DB
func updateTimeProp(timeEntryId uint, propName string, value interface{}) {
	dbConn := ConnectDB()
	query := `UPDATE tbl_timeentry SET updated = NOW(), ` + propName + ` = $1 WHERE id = $2`
	_, err := dbConn.Exec(query, value, timeEntryId)
	helpers.HandleError(err)
}

// Deletes the selected time entry
func DeleteTimeEntry(userId, timeEntryId uint) error {
	// Check that the time entry and user ids are valid and that the user
	// can delete this time entry and return an error where necessary
	if err := checkUserIdAndTimeEntryIdValue(userId, timeEntryId); err != nil {
		return err
	}

	// Connect to the database and remove delete the selected time entry
	dbConn := ConnectDB()
	_, err := dbConn.Exec(`call sp_time_delete($1)`, timeEntryId)

	// Return the error received from the database
	// if nil, deletion was successful
	return err
}

// Checks if the user and time entry ids are valid
func checkUserIdAndTimeEntryIdValue(userId, timeEntryId uint) error {
	// Return invalid time entry id error if time entry is 0
	if timeEntryId == 0 {
		return errors.New("invalid time entry Id")
	}

	// Return an invalid user error if the user cannot be found
	if userId == 0 {
		return errors.New("invalid user id")
	}

	// Return an access denied error if user is trying to
	// amend a time entry they do not have permissions for
	if !canAmendTimeEntry(userId, timeEntryId) {
		return errors.New("access denied: cannot amend the selected time entry")
	}

	// all the above is fine, so return nil to indicate this
	return nil
}

// Checks if the user can amend or delete the selected time entry
func canAmendTimeEntry(userId, timeEntryId uint) bool {
	dbConn := ConnectDB()

	// Get the user ID for the selected time entry from the database
	var timeUserId uint = 0
	row := dbConn.QueryRow(`SELECT userId FROM tbl_timeentry WHERE id = $1`, timeEntryId)
	err := row.Scan(&timeUserId)

	// handle any errors returned from the database
	helpers.HandleError(err)

	// Check that the user trying to amend a time entry
	// is the owner/creator of the time entry
	return timeUserId == userId
}
