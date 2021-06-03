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
	id            uuid.UUID
	name          string
	email         string
	organisations *[]Organisation
	created       time.Time
	updated       time.Time
	githubUserId  string
	avatar        string
	apiClients    *[]ApiClient
}

type Organisation struct {
	id          uuid.UUID
	name        string
	description string
	avatar      string
}

type RepoItem struct {
	id           uint
	timeEntryId  uint
	itemIdSource string
	itemType     string
	source       string
	created      time.Time
	repoName     string
	description  string
}

type Tag struct {
	id   uint
	name string
}

type TimeEntry struct {
	id             uint
	userId         uuid.UUID
	organisationId uuid.UUID
	comments       string
	created        time.Time
	updated        time.Time
	value          float32
	valueType      string
	tags           []Tag
	repoItems      []RepoItem
}
