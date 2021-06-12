package usersdb

import (
	"database/sql"
	"timetracker/helpers"
	"timetracker/models"

	"github.com/google/go-github/v35/github"
)

// Checks if the githubUserId exists in the database
func GitHubUserExists(githubUserId string) bool {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	query := `SELECT COUNT(*) AS count FROM tbl_user WHERE githubUserId = $1`
	row := dbConn.QueryRow(query, githubUserId)

	var count uint
	err := row.Scan(&count)
	helpers.HandleError(err)

	return count > 0
}

// Creates a new user from the github user details
func CreateUser(u github.User) models.User {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	name := u.Name
	if name == nil {
		name = u.Login
	}

	query := `
		INSERT INTO tbl_user(name, email, githubUserId, avatar)
		VALUES($1, $2, $3, $4)
		RETURNING id, name, coalesce(email, '') AS email, created, updated, githubUserId, avatar
	`
	row := dbConn.QueryRow(query, name, u.Email, u.Login, u.AvatarURL)

	return readUserFromSqlRow(row)
}

// Gets a user's details from the db by the github login id
func GetUserByGitHubLogin(githubUserId string) models.User {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	query := `
		SELECT id, name, coalesce(email, '') AS email, created, updated, githubUserId, avatar 
		FROM tbl_user 
		WHERE githubUserId = $1
	`
	row := dbConn.QueryRow(query, githubUserId)

	return readUserFromSqlRow(row)
}

// Reads a user from a sql row
func readUserFromSqlRow(row *sql.Row) models.User {
	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Created, &user.Updated, &user.GithubUserId, &user.Avatar)
	helpers.HandleError(err)

	user.ApiClients = []models.ApiClient{}
	user.Organisations = []models.Organisation{}

	return user
}
