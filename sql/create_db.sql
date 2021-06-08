--CREATE DATABASE timetracker;

SHOW SEARCH_PATH;

SET SEARCH_PATH TO timetracker;

DROP TABLE IF EXISTS public.UserOrgLink;
DROP TABLE IF EXISTS public.TimeEntryTagLink;
DROP TABLE IF EXISTS public.Tag;
DROP TABLE IF EXISTS public.RepoItem;
DROP TABLE IF EXISTS public.TimeEntry;
DROP TABLE IF EXISTS public.ApiClient;
DROP TABLE IF EXISTS public.Organisation;
DROP TABLE IF EXISTS public.User;

CREATE TABLE public.User(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(400),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    githubUserId VARCHAR(100),
    avatar TEXT
);

CREATE INDEX IX_user_githubUserId ON public.User (githubuserid);
CREATE INDEX IX_user_email ON public.User (email);

CREATE TABLE public.Organisation(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description varchar(500),
    avatar TEXT,
    source VARCHAR(50),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE public.UserOrgLink(
    organisationId INT,
    userId INT
);

ALTER TABLE public.UserOrgLink
    ADD CONSTRAINT FK_UserOrgLink_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES public.Organisation (id)
        ON DELETE CASCADE;

ALTER TABLE public.UserOrgLink
    ADD CONSTRAINT FK_UserOrgLink_UserId 
        FOREIGN KEY (userId)
        REFERENCES public.User (id)
        ON DELETE CASCADE;

CREATE TABLE public.ApiClient(
    id SERIAL PRIMARY KEY,
    clientId UUID,
    userId INT,
    secretKey UUID,
    appName VARCHAR(50) NOT NULL,
    description VARCHAR(500),
    validTill TIMESTAMP,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IX_apiclient_clientId ON ApiClient (clientId);

ALTER TABLE public.ApiClient
    ADD CONSTRAINT FK_ApiClient_UserId
    FOREIGN KEY (userId)
    REFERENCES public.User (id)
    ON DELETE CASCADE;

CREATE TABLE public.TimeEntry(
    id BIGSERIAL PRIMARY KEY,
    comments VARCHAR(200),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    value NUMERIC NOT NULL,
    valueType varchar(20) NOT NULL,
    userId INT,
    organisationId INT
);

CREATE INDEX IX_timeentry_created ON public.TimeEntry (created);
CREATE INDEX IX_timeentry_updated ON public.TimeEntry (updated);

ALTER TABLE public.TimeEntry
    ADD CONSTRAINT FK_TimeEntry_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES public.Organisation (id);

ALTER TABLE public.TimeEntry
    ADD CONSTRAINT FK_TimeEntry_UserId 
        FOREIGN KEY (userId)
        REFERENCES public.User (id);

CREATE TABLE public.RepoItem(
    id BIGSERIAL PRIMARY KEY,
    itemIdSource VARCHAR(200) NOT NULL,
    itemType VARCHAR(50) NOT NULL,
    source VARCHAR(50) NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    description VARCHAR(2000),
    timeEntryId BIGINT
);

CREATE INDEX IX_repoitem_source ON RepoItem (source);

ALTER TABLE public.RepoItem
    ADD CONSTRAINT FK_RepoItem_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES public.TimeEntry (id)
        ON DELETE CASCADE;

CREATE TABLE public.Tag(
    id SERIAL PRIMARY KEY,
    name varchar(50)
);

CREATE TABLE public.TimeEntryTagLink(
    tagId INT,
    timeEntryId BIGINT
);

ALTER TABLE public.TimeEntryTagLink
    ADD CONSTRAINT FK_TimeEntryTagLink_TagId 
        FOREIGN KEY (tagId)
        REFERENCES public.Tag (id)
        ON DELETE CASCADE;

ALTER TABLE public.TimeEntryTagLink
    ADD CONSTRAINT FK_TimeEntryTagLink_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES public.TimeEntry (id)
        ON DELETE CASCADE;