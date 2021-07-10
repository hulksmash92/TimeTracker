package db

import (
	"timetracker/helpers"
	"timetracker/models"
)

// Gets all tags from the DB
func GetTags() []models.Tag {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	rows, err := dbConn.Query("SELECT * FROM tbl_tag")
	helpers.HandleError(err)
	defer rows.Close()

	tags := []models.Tag{}

	for rows.Next() {
		var t models.Tag
		rows.Scan(&t.Id, &t.Name)
		tags = append(tags, t)
	}

	return tags
}
