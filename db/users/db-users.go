package usersdb

import (
	"timetracker/helpers"
	"timetracker/models"
)

func GitHubUserExists(githubUserId string) bool {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	query := `SELECT COUNT(*) AS count FROM tbl_users WHERE githubUserId = $1`
	rows, err := dbConn.Query(query, githubUserId)
	defer rows.Close()
	helpers.HandleError(err)

	for rows.Next() {
		var count uint
		err := rows.Scan(&count)
		helpers.HandleError(err)

		if count > 0 {
			return true
		}
	}

	return false
}

func CreateUser(user models.User) {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

}
