package timedb

import (
	"time"
	"timetracker/helpers"
	"timetracker/models"
)

func GetTimeEntries(userId uint, dateFrom time.Time, dateTo time.Time) []models.TimeEntry {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	query := `
		SELECT t.id, t.created, t.updated, t.value, t.valueType,
			coalesce(t.comments, '') AS comments,
			coalesce(t.userId, 0) AS userId,
			coalesce(u.name, '') AS username,
			coalesce(u.avatar, '') AS userAvatar,
			coalesce(t.organisationId, 0) AS organisationId,
			coalesce(o.name, '') AS organisation,
			coalesce(o.avatar, '') AS organisationAvatar
		FROM tbl_timeentry AS t
		LEFT JOIN tbl_user AS u ON u.id = t.userId
		LEFT JOIN tbl_organisation AS o ON o.id = t.organisationId
		WHERE t.userId = $1 AND t.created >= $2 AND t.created <= dateTo
	`
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
