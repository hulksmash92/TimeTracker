package db

import (
	"database/sql"
	"timetracker/helpers"
	"timetracker/models"

	"github.com/google/go-github/v35/github"
)

// Gets the user id for the selected github user
func GetUserId(githubUserId string) uint {
	dbConn := helpers.ConnectDB()
	defer dbConn.Close()

	query := `SELECT id FROM tbl_user WHERE githubUserId = $1`
	row := dbConn.QueryRow(query, githubUserId)

	var id uint
	err := row.Scan(&id)
	helpers.HandleError(err)

	return id
}

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
	user := readUserFromSqlRow(row)

	query = `
		SELECT clientId, appName, description, validTill, created, updated 
		FROM tbl_apiclient 
		WHERE userId = $1
	`
	clientRows, err := dbConn.Query(query, user.Id)
	helpers.HandleError(err)
	defer clientRows.Close()

	for clientRows.Next() {
		var a models.ApiClient
		clientRows.Scan(&a.ClientId, &a.AppName, &a.Description, &a.ValidTill, &a.Created, &a.Updated)
		user.ApiClients = append(user.ApiClients, a)
	}

	query = `
		SELECT o.id, o.name, o.description, o.avatar, o.source, o.updated
		FROM tbl_userorglink AS uol
		INNER JOIN tbl_organisation AS o ON uol.organisationid = o.id
		WHERE uol.userid = $1
	`
	orgRows, err := dbConn.Query(query, user.Id)
	helpers.HandleError(err)
	defer orgRows.Close()

	for orgRows.Next() {
		var o models.Organisation
		orgRows.Scan(&o.Id, &o.Name, &o.Description, &o.Avatar, &o.Source, &o.Created, &o.Updated)
		user.Organisations = append(user.Organisations, o)
	}

	return user
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
