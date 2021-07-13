package models

import (
	"time"

	"github.com/google/uuid"
)

// Details of client application credentials a user has
// to query the web API programatically
type ApiClient struct {
	// The client's unique ID
	ClientId uuid.UUID `json:"clientId"`

	// Name of the client application
	AppName string `json:"appName"`

	// Description of the application
	Description string `json:"description"`

	// When the credentials are valid until
	ValidTill time.Time `json:"validTill"`

	// When the entry was created
	Created time.Time `json:"created"`

	// When the entry was last amended
	Updated time.Time `json:"updated"`
}

// Details of an application user
type User struct {
	// Unique ID of the user in the database
	Id uint `json:"id"`

	// Display name of the user
	Name string `json:"name"`

	// User's email address
	Email string `json:"email"`

	// Date & time the user was created
	Created time.Time `json:"created"`

	// Date & time hen the user was last amended
	Updated time.Time `json:"updated"`

	// User's GitHub login name
	GithubUserId string `json:"githubUserId"`

	// URL of the user's avatar, will be pulled from GitHub
	// when they sign in for the first time
	Avatar string `json:"avatar"`

	// Details of any organisations a user is linked to
	Organisations []Organisation `json:"organisations"`

	// Details of any client applications the user owns
	ApiClients []ApiClient `json:"apiClients"`
}

// Details of an organisation
type Organisation struct {
	// Unique ID in the DB
	Id uint `json:"id"`

	// Organisation display name
	Name string `json:"name"`

	// Further details of the organisation
	Description string `json:"description"`

	// URL or base64 encoded binary data of the avatar/photo
	Avatar string `json:"avatar"`

	// Where the organisation originated, i.e. GitHub
	Source string `json:"source"`

	// Date & time the organisation was created
	Created time.Time `json:"created"`

	// Date & time the organisation was last amended
	Updated time.Time `json:"updated"`
}

// Details of an item from a Git repository
// like commits or branches
type RepoItem struct {
	// ID of the item in the DB
	Id uint `json:"id"`

	// ID of the item from the originated source
	ItemIdSource string `json:"itemIdSource"`

	// Type of the item, i.e. commit
	ItemType string `json:"itemType"`

	// Where the item came from, i.e. GitHub
	Source string `json:"source"`

	// Date & time when this item was created in the DB
	Created time.Time `json:"created"`

	// Date & time when this item was amended in the DB
	Updated time.Time `json:"updated"`

	// Name of the repository when the item came from
	RepoName string `json:"repoName"`

	// May be the commit's message or message of a branches last commit
	Description string `json:"description"`
}

// Details of a tag to categorise TimeEntry
type Tag struct {
	// Unique ID of the tag in DB
	Id uint `json:"id"`

	// Display name for the tag
	Name string `json:"name"`
}

// Represents a trimmed down structure of time entry
// This can be the linked user or organisation
type OwnerTrimmed struct {
	// ID of the user or organisation in the DB
	Id uint `json:"id"`

	// Name of the user or organisation
	Name string `json:"name"`

	// URL of the user or organisation's avatar
	Avatar string `json:"avatar"`
}

// Represents a created time entry object in the database
type TimeEntry struct {
	// Unique ID of the time entry in the database
	Id uint `json:"id"`

	// User who created the time entry
	User OwnerTrimmed `json:"user"`

	// Organisation that the work this time entry was completed for
	// Users can be members of multiple organisations
	Organisation OwnerTrimmed `json:"organisation"`

	// Descriptive comments about this time entry, i.e. what was done & why
	Comments string `json:"comments"`

	// Date & time when the time entry was created
	Created time.Time `json:"created"`

	// Date & time when the time entry was last amended
	Updated time.Time `json:"updated"`

	// Number of hours, minutes or units recorded
	Value float32 `json:"value"`

	// Type of the above value, i.e. hours, minutes or time units
	ValueType string `json:"valueType"`

	// Any linked tags that categorise this time entry
	// for example: Feature, Bug etc.
	Tags []Tag `json:"tags,omitempty"`

	// Any Git repository items linked to this entry like branches or commits
	RepoItems []RepoItem `json:"repoItems,omitempty"`
}
