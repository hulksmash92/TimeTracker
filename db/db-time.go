package db

import (
	"errors"
	"time"
	"timetracker/helpers"
	"timetracker/models"
)

type UpdatedTimeEntry struct {
	Comments  *string
	Value     *float32
	ValueType *string
	Tags      *[]models.Tag
	RepoItems *[]models.RepoItem
}

// Gets all time entries for a user and the given date range
func GetTimeEntries(userId uint, dateFrom time.Time, dateTo time.Time) []models.TimeEntry {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	query := `call sp_time_get($1, $2, $3)`
	rows, err := dbConn.Query(query, userId, dateFrom, dateTo)
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

	query := `call sp_time_get_by_id($1)`
	row := dbConn.QueryRow(query, id)

	var t models.TimeEntry
	var u models.OwnerTrimmed
	var o models.OwnerTrimmed
	err := row.Scan(&t.Id, &t.Created, &t.Updated, &t.Value, &t.ValueType, &t.Comments, &u.Id, &u.Name, &u.Avatar, &o.Id, &o.Name, &o.Avatar)
	helpers.HandleError(err)
	t.User = u
	t.Organisation = o
	t.RepoItems = []models.RepoItem{}
	t.Tags = []models.Tag{}

	query = `call sp_time_tags($1)`
	rows, err := dbConn.Query(query, id)
	helpers.HandleError(err)
	defer rows.Close()

	for rows.Next() {
		var tg models.Tag
		rows.Scan(&tg)
		t.Tags = append(t.Tags, tg)
	}

	query = `call sp_time_repoitems($1)`
	repoRows, err := dbConn.Query(query, id)
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

	query := `call sp_time_insert($1, $2, $3, $4, $5)`
	row := dbConn.QueryRow(query, t.User.Id, t.Organisation.Id, t.Comments, t.Value, t.ValueType)
	err := row.Scan(&t.Id)
	helpers.HandleError(err)

	for _, tag := range t.Tags {
		query = `call sp_time_tags_insert($1, $2, $3)`
		dbConn.Exec(query, t.Id, tag.Name, tag.Id)
	}

	for _, r := range t.RepoItems {
		query = `call sp_time_repoitem_insert($1, $2, $3, $4, $5, $6, $7)`
		dbConn.Exec(query, t.Id, r.Created, r.ItemIdSource, r.ItemType, r.Source, r.RepoName, r.Description)
	}

	return t.Id
}

// Updates an existing time entry
func UpdateTimeEntry(userId uint, timeEntryId uint, vals UpdatedTimeEntry) error {
	if timeEntryId == 0 {
		return errors.New("Invalid time entry Id")
	}
	if userId == 0 {
		return errors.New("Invalid user id")
	}
	if !canAmendTimeEntry(userId, timeEntryId) {
		return errors.New("Access denied: cannot amend the selected time entry")
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
		query = `DELETE FROM tbl_timeentrytaglinks WHERE timeentryid = $1`
		_, err := dbConn.Exec(query, vals.ValueType, timeEntryId)
		helpers.HandleError(err)

		for _, tag := range *vals.Tags {
			query = `call sp_time_tags_insert($1, $2, $3)`
			dbConn.Exec(query, timeEntryId, tag.Name, tag.Id)
		}
	}

	if vals.RepoItems != nil {
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
	if timeEntryId == 0 {
		return errors.New("Invalid time entry Id")
	}
	if userId == 0 {
		return errors.New("Invalid user id")
	}
	if !canAmendTimeEntry(userId, timeEntryId) {
		return errors.New("Access denied: cannot amend the selected time entry")
	}

	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	query := `call sp_time_delete($1)`
	_, err := dbConn.Exec(query, timeEntryId)
	helpers.HandleError(err)

	return nil
}

func canAmendTimeEntry(userId uint, timeEntryId uint) bool {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	var timeUserId uint
	timeEntryId = 0

	return timeUserId == userId
}
