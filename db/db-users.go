package db

import (
	"database/sql"
	"timetracker/helpers"
	"timetracker/models"

	"github.com/google/go-github/v35/github"
)

// Gets the user id for the selected github user
func GetUserId(githubUserId string) uint {
	dbConn := ConnectDB()
	query := `SELECT id FROM tbl_user WHERE githubUserId = $1`
	row := dbConn.QueryRow(query, githubUserId)

	var id uint
	err := row.Scan(&id)
	helpers.HandleError(err)

	return id
}

// Checks if the githubUserId exists in the database
func GitHubUserExists(githubUserId string) bool {
	// Connect to the database
	dbConn := ConnectDB()

	// Query the DB using a parameterised SQL statement to get a single value
	query := `SELECT COUNT(*) AS count FROM tbl_user WHERE githubUserId = $1`
	row := dbConn.QueryRow(query, githubUserId)

	var count uint
	err := row.Scan(&count)

	// Handle any errors from parsing the database resul
	helpers.HandleError(err)

	// if user exists, count will be 1
	return count == 1
}

// Creates a new user from the github user details
func CreateUser(u github.User) models.User {
	dbConn := ConnectDB()

	// If name field on github user is blank, set to the login ID
	name := u.Name
	if name == nil {
		name = u.Login
	}

	// Create a parameterised SQL statement for creating a new user
	// that returns a row of user data
	query := `
		INSERT INTO tbl_user(name, email, githubUserId, avatar)
		VALUES($1, $2, $3, $4)
		RETURNING id, name, coalesce(email, '') AS email, created, updated, githubUserId, avatar
	`

	// Get a single row of data and parse this to a User object
	row := dbConn.QueryRow(query, name, u.Email, u.Login, u.AvatarURL)
	return readUserFromSqlRow(row)
}

// Gets a user's details from the db by the github login id
func GetUserByGitHubLogin(githubUserId string) models.User {
	dbConn := ConnectDB()

	// Create a query to get the existing user by their GitHub user ID
	query := `
		SELECT id, name, coalesce(email, '') AS email, created, updated, githubUserId, avatar 
		FROM tbl_user 
		WHERE githubUserId = $1
	`

	// Use the above query to get the user row
	row := dbConn.QueryRow(query, githubUserId)
	return readUserFromSqlRow(row)

	//user := readUserFromSqlRow(row)

	// query = `
	// 	SELECT clientId, appName, description, validTill, created, updated
	// 	FROM tbl_apiclient
	// 	WHERE userId = $1
	// `
	// clientRows, err := dbConn.Query(query, user.Id)
	// helpers.HandleError(err)
	// defer clientRows.Close()

	// for clientRows.Next() {
	// 	var a models.ApiClient
	// 	clientRows.Scan(&a.ClientId, &a.AppName, &a.Description, &a.ValidTill, &a.Created, &a.Updated)
	// 	user.ApiClients = append(user.ApiClients, a)
	// }

	// query = `
	// 	SELECT o.id, o.name, o.description, o.avatar, o.source, o.updated
	// 	FROM tbl_userorglink AS uol
	// 	INNER JOIN tbl_organisation AS o ON uol.organisationid = o.id
	// 	WHERE uol.userid = $1
	// `
	// orgRows, err := dbConn.Query(query, user.Id)
	// helpers.HandleError(err)
	// defer orgRows.Close()

	// for orgRows.Next() {
	// 	var o models.Organisation
	// 	orgRows.Scan(&o.Id, &o.Name, &o.Description, &o.Avatar, &o.Source, &o.Created, &o.Updated)
	// 	user.Organisations = append(user.Organisations, o)
	// }

	// Return the parsed user
	//return user
}

// Reads a user from a sql row
func readUserFromSqlRow(row *sql.Row) models.User {
	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Created, &user.Updated, &user.GithubUserId, &user.Avatar)
	helpers.HandleError(err)

	// Leave these as empty arrays for now, will be populated when orgs and api clients are enabled
	user.ApiClients = []models.ApiClient{}
	user.Organisations = []models.Organisation{}

	return user
}

// Updates the user's profile
func UpdateUserProfile(userId uint, name, email *string) {
	newName := ""
	newEmail := ""

	if name != nil {
		newName = *name
	}
	if email != nil {
		newEmail = *email
	}

	dbConn := ConnectDB()
	_, err := dbConn.Exec("call sp_update_user($1, $2, $3)", userId, newName, newEmail)
	helpers.HandleError(err)
}

// Deletes the selected user
func DeleteUser(userId uint) {
	dbConn := ConnectDB()
	_, err := dbConn.Exec("DELETE FROM tbl_user WHERE id = $1", userId)
	helpers.HandleError(err)
}
