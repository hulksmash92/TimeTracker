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
func GetTimeEntries(userId uint, dateFrom time.Time, dateTo time.Time) []models.TimeEntry {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	rows, err := dbConn.Query(`call sp_time_get($1, $2, $3)`, userId, dateFrom, dateTo)
	helpers.HandleError(err)
	defer rows.Close()
	time := []models.TimeEntry{}

	for rows.Next() {
		var t models.TimeEntry
		var u models.OwnerTrimmed
		var o models.OwnerTrimmed
		rows.Scan(&t.Id, &t.Created, &t.Updated, &t.Value, &t.ValueType, &t.Comments, &u.Id, &u.Name, &u.Avatar, &o.Id, &o.Name, &o.Avatar)
		t.User = u
		t.Organisation = o
		time = append(time, t)
	}

	return time
}

// Gets the time entry with a given id
func GetTimeEntry(id uint) models.TimeEntry {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	row := dbConn.QueryRow(`call sp_time_get_by_id($1)`, id)
	var t models.TimeEntry
	var u models.OwnerTrimmed
	var o models.OwnerTrimmed
	err := row.Scan(&t.Id, &t.Created, &t.Updated, &t.Value, &t.ValueType, &t.Comments, &u.Id, &u.Name, &u.Avatar, &o.Id, &o.Name, &o.Avatar)
	helpers.HandleError(err)
	t.User = u
	t.Organisation = o
	t.RepoItems = []models.RepoItem{}
	t.Tags = []models.Tag{}

	// Get the tags for the time entry
	rows, err := dbConn.Query(`call sp_time_tags($1)`, id)
	helpers.HandleError(err)
	defer rows.Close()
	for rows.Next() {
		var tg models.Tag
		rows.Scan(&tg)
		t.Tags = append(t.Tags, tg)
	}

	// Get the repo items added to the time entry
	repoRows, err := dbConn.Query(`call sp_time_repoitems($1)`, id)
	helpers.HandleError(err)
	defer repoRows.Close()
	for repoRows.Next() {
		var r models.RepoItem
		repoRows.Scan(&r.Id, &r.Created, &r.Updated, &r.ItemIdSource, &r.ItemType, &r.Source, &r.RepoName, &r.Description)
		t.RepoItems = append(t.RepoItems, r)
	}

	return t
}

// Creates a new time entry in the database
func CreateTimeEntry(t models.TimeEntry) uint {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	row := dbConn.QueryRow(`call sp_time_insert($1, $2, $3, $4, $5)`, t.User.Id, t.Organisation.Id, t.Comments, t.Value, t.ValueType)
	err := row.Scan(&t.Id)
	helpers.HandleError(err)

	for _, tag := range t.Tags {
		dbConn.Exec(`call sp_time_tags_insert($1, $2, $3)`, t.Id, tag.Name, tag.Id)
	}

	for _, r := range t.RepoItems {
		query := `call sp_time_repoitem_insert($1, $2, $3, $4, $5, $6, $7)`
		dbConn.Exec(query, t.Id, r.Created, r.ItemIdSource, r.ItemType, r.Source, r.RepoName, r.Description)
	}

	return t.Id
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
				dbConn.Exec(`DELETE FROM tbl_repoitem WHERE id = $1`, id)
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
