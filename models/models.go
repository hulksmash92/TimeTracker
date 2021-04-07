package models

import (
	"time"

	"github.com/google/uuid"
)

type ApiClient struct {
	clientId    uuid.UUID
	secretKey   uuid.UUID
	appName     string
	description string
	validTill   time.Time
	updated     time.Time
	userId      uuid.UUID
}

type User struct {
	userId            uuid.UUID
	name              string
	email             string
	encryptedPassword string
	organisation      string
	created           time.Time
	updated           time.Time
	githubUserId      string
	avatar            string
	apiClients        []ApiClient
}

type RepoItem struct {
	itemId      uuid.UUID
	timeEntryId uint
	itemType    string
	source      string
	created     time.Time
	repooName   string
	description string
}

type TimeEntry struct {
	entryId   uint
	userId    uuid.UUID
	created   time.Time
	updated   uuid.UUID
	comments  string
	tags      []string
	repoItems []RepoItem
	value     float32
	valueType string
}
