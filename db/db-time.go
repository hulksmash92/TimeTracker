package db

import (
	"errors"
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

// Gets all time entries for a user and the given date range
// TODO: Add server side pagination and sorting to improve speeds as data scales
func GetTimeEntries(userId uint, dateFrom time.Time, dateTo time.Time) []models.TimeEntry {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	// Create our parameterised SQL query using our vw_time_entries view
	query := `
		SELECT id, created, updated, value, valueType, comments, 
			userId, username, userAvatar, organisationId,
			organisation, organisationAvatar
		FROM vw_time_entries AS t
		WHERE t.userId = $1 AND t.created >= $2 AND t.created <= $3
		ORDER BY created DESC;
	`
	// Query the db with the above statement and handle any returned errors
	rows, err := dbConn.Query(query, userId, dateFrom, dateTo)
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

	// return an parsed rows
	return time
}

// Gets the time entry with a given id
func GetTimeEntry(id uint) models.TimeEntry {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	// Create our parameterised SQL query using our vw_time_entries view
	query := `
		SELECT id, created, updated, value, valueType, comments, 
			userId, username, userAvatar, organisationId,
			organisation, organisationAvatar
		FROM vw_time_entries AS t
		WHERE t.id = $1;
	`

	// Get a single row of data for the above query
	row := dbConn.QueryRow(query, id)

	var timeEntry models.TimeEntry
	var user models.OwnerTrimmed
	var org models.OwnerTrimmed

	// Parse the row to the above data model
	// if no time entries with the id where found,
	// an error will be thrown to indicate this
	err := row.Scan(
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

	// Handle any returned errors from row.Scan
	helpers.HandleError(err)

	// Set the user and org properties to the parsed values below
	timeEntry.User = user
	timeEntry.Organisation = org

	// Set the RepoItems and Tags properties to empty arrays
	timeEntry.RepoItems = []models.RepoItem{}
	timeEntry.Tags = []models.Tag{}

	// Get the tags for the time entry and handle any returned errors
	tagRows, err := dbConn.Query(`SELECT id, name FROM vw_time_entry_tags WHERE timeentryid = $1`, id)
	helpers.HandleError(err)

	// Defer closing of the tagRows cursor till the function returns
	defer tagRows.Close()

	// loop through the rows returned and parse them to an instance of Tag
	// and append each value to timeEntry.Tags
	for tagRows.Next() {
		var tag models.Tag
		tagRows.Scan(&tag.Id, &tag.Name)
		timeEntry.Tags = append(timeEntry.Tags, tag)
	}

	// linked repo items prepared statement
	query = `
		SELECT id, created, updated, itemIdSource, itemType, source, repoName, description
		FROM vw_repo_items 
		WHERE timeentryid = $1;
	`

	// Get the repo items added to the time entry and handle any errors
	repoRows, err := dbConn.Query(query, id)
	helpers.HandleError(err)

	// Defer closing of the repoRows cursor till the function returns
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

		timeEntry.RepoItems = append(timeEntry.RepoItems, repoItem)
	}

	// Return the parsed time entry
	return timeEntry
}

// Creates a new time entry in the database from the passed in models.TimeEntry
// and returns the id of the new time entry
func CreateTimeEntry(newEntry models.TimeEntry) uint {
	// Connect to the DB and defer closing the connection until
	// this function is finished
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

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

	// loop through the tags added to the entry and add the new tags
	// and/or the link between the tags and time entry using our
	// sp_time_tags_insert stored procedure
	for _, tag := range newEntry.Tags {
		dbConn.Exec(`call sp_time_tags_insert($1, $2, $3)`,
			newEntry.Id,
			tag.Name,
			tag.Id)
	}

	// Loop through the linked repository items to the database with a link to this
	// new time entry, using the sp_time_repoitem_insert stored procedure
	for _, r := range newEntry.RepoItems {
		dbConn.Exec(`call sp_time_repoitem_insert($1, $2, $3, $4, $5, $6, $7)`,
			newEntry.Id,
			r.Created,
			r.ItemIdSource,
			r.ItemType,
			r.Source,
			r.RepoName,
			r.Description)
	}

	// Return the time entries new id
	return newEntry.Id
}

// Updates an existing time entry
func UpdateTimeEntry(userId uint, timeEntryId uint, vals UpdatedTimeEntry) error {
	if err := checkUserIdAndTimeEntryIdValue(userId, timeEntryId); err != nil {
		return err
	}

	dbConn := helpers.ConnectDB()
	defer dbConn.Close()
	var query string

	if vals.Comments != nil {
		query = `UPDATE tbl_timeentry SET updated = NOW(), comments = $1 WHERE id = $2`
		_, err := dbConn.Exec(query, vals.Comments, timeEntryId)
		helpers.HandleError(err)
	}
	if vals.Value != nil {
		query = `UPDATE tbl_timeentry SET updated = NOW(), value = $1 WHERE id = $2`
		_, err := dbConn.Exec(query, vals.Value, timeEntryId)
		helpers.HandleError(err)
	}
	if vals.ValueType != nil {
		query = `UPDATE tbl_timeentry SET updated = NOW(), valueType = $1 WHERE id = $2`
		_, err := dbConn.Exec(query, vals.ValueType, timeEntryId)
		helpers.HandleError(err)
	}
	if vals.Tags != nil {
		_, err := dbConn.Exec(`DELETE FROM tbl_timeentrytaglinks WHERE timeentryid = $1`, vals.ValueType, timeEntryId)
		helpers.HandleError(err)

		for _, tag := range *vals.Tags {
			dbConn.Exec(`call sp_time_tags_insert($1, $2, $3)`, timeEntryId, tag.Name, tag.Id)
		}
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
		for _, r := range *vals.RepoItems {
			if r.Id == 0 {
				query = `call sp_time_repoitem_insert($1, $2, $3, $4, $5, $6, $7)`
				dbConn.Exec(query, timeEntryId, r.Created, r.ItemIdSource, r.ItemType, r.Source, r.RepoName, r.Description)
			}
		}
	}

	return nil
}

// Deletes the selected time entry
func DeleteTimeEntry(userId uint, timeEntryId uint) error {
	if err := checkUserIdAndTimeEntryIdValue(userId, timeEntryId); err != nil {
		return err
	}

	dbConn := helpers.ConnectDB()
	defer dbConn.Close()
	_, err := dbConn.Exec(`call sp_time_delete($1)`, timeEntryId)
	helpers.HandleError(err)

	return nil
}

// Checks if the user and time entry ids are valid
func checkUserIdAndTimeEntryIdValue(userId uint, timeEntryId uint) error {
	if timeEntryId == 0 {
		return errors.New("Invalid time entry Id")
	}
	if userId == 0 {
		return errors.New("Invalid user id")
	}
	if !canAmendTimeEntry(userId, timeEntryId) {
		return errors.New("Access denied: cannot amend the selected time entry")
	}
	return nil
}

// Checks if the user can amend the selected time entry
func canAmendTimeEntry(userId uint, timeEntryId uint) bool {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	var timeUserId uint = 0
	row := dbConn.QueryRow(`SELECT userId FROM tbl_timeentry WHERE id = $1`, timeEntryId)
	err := row.Scan(&timeUserId)
	helpers.HandleError(err)

	return timeUserId == userId
}
