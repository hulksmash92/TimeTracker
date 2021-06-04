--CREATE DATABASE timetracker;

CREATE TABLE User (
    id UUID PRIMARY KEY NOT NULL
    name VARCHAR(200) NOT NULL
    email VARCHAR(400)
    created TIMESTAMP NOT NULL DEFAULT NOW()
    updated TIMESTAMP NOT NULL DEFAULT NOW()
    githubUserId VARCHAR(100)
    avatar TEXT
);

CREATE TABLE Organisation (
    id UUID PRIMARY KEY NOT NULL
    name VARCHAR(200) NOT NULL
    decription varchar(500)
    avatar TEXT
    source VARCHAR(50)
    created TIMESTAMP NOT NULL DEFAULT NOW()
    updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE UserOrgLink (
    organisationId UUID
    userId UUID
);

ALTER TABLE UserOrgLink 
    ADD CONSTRAINT FK_UserOrgLink_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES Organisation (id)
        ON DELETE CASCADE;

ALTER TABLE UserOrgLink 
    ADD CONSTRAINT FK_UserOrgLink_UserId 
        FOREIGN KEY (userId)
        REFERENCES User (id)
        ON DELETE CASCADE;

CREATE TABLE ApiClient (
    clientId UUID PRIMARY KEY NOT NULL
    userId UUID
    secretKey UUID
    appName VARCHAR(50) NOT NULL
    description VARCHAR(500)
    validTill TIMESTAMP
    created TIMESTAMP NOT NULL DEFAULT NOW()
    updated TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE ApiClient
    ADD CONSTRAINT FK_ApiClient_UserId
    REFERENCES User (id)
    ON DELETE CASCADE;

CREATE TABLE TimeEntry (
    id INT PRIMARY KEY NOT NULL IDENTITY(1, 1)
    comments VARCHAR(200)
    created TIMESTAMP NOT NULL DEFAULT NOW()
    updated TIMESTAMP NOT NULL DEFAULT NOW()
    value NUMERIC NOT NULL
    valueType varchar(20) NOT NULL
    userId UUID
    organisationId UUID
);

ALTER TABLE TimeEntry 
    ADD CONSTRAINT FK_TimeEntry_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES Organisation (id);

ALTER TABLE TimeEntry 
    ADD CONSTRAINT FK_TimeEntry_UserId 
        FOREIGN KEY (userId)
        REFERENCES User (id);

CREATE TABLE RepoItem (
    id INT PRIMARY KEY NOT NULL IDENTITY(1, 1),
    itemIdSource VARCHAR(200) NOT NULL
    itemType VARCHAR(50) NOT NULL
    source VARCHAR(50) NOT NULL
    created TIMESTAMP NOT NULL
    updated TIMESTAMP NOT NULL DEFAULT NOW()
    description VARCHAR(2000)
    timeEntryId INT
);

ALTER TABLE RepoItem 
    ADD CONSTRAINT FK_RepoItem_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES TimeEntry (id)
        ON DELETE CASCADE;

CREATE TABLE Tag (
    id INT PRIMARY KEY NOT NULL IDENTITY(1, 1)
    name varchar(50)
);

CREATE TimeEntryTagLink (
    tagId INT
    timeEntryId INT
);

ALTER TABLE TimeEntryTagLink 
    ADD CONSTRAINT FK_TimeEntryTagLink_TagId 
        FOREIGN KEY (tagId)
        REFERENCES Tag (id)
        ON DELETE CASCADE;

ALTER TABLE TimeEntryTagLink 
    ADD CONSTRAINT FK_TimeEntryTagLink_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES TimeEntry (id)
        ON DELETE CASCADE;