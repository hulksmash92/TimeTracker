--CREATE DATABASE timetracker;

SET SEARCH_PATH TO timetracker;

SET SEARCH_PATH TO public;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS public.tbl_UserOrgLink;
DROP TABLE IF EXISTS public.tbl_TimeEntryTagLink;
DROP TABLE IF EXISTS public.tbl_Tag;
DROP TABLE IF EXISTS public.tbl_RepoItem;
DROP TABLE IF EXISTS public.tbl_TimeEntry;
DROP TABLE IF EXISTS public.tbl_ApiClient;
DROP TABLE IF EXISTS public.tbl_Organisation;
DROP TABLE IF EXISTS public.tbl_User;

CREATE TABLE public.tbl_User(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(400),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    githubUserId VARCHAR(100),
    avatar TEXT
);

CREATE INDEX IX_user_githubUserId ON public.tbl_User (githubuserid);
CREATE INDEX IX_user_email ON public.tbl_User (email);

CREATE TABLE public.tbl_Organisation(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description varchar(500),
    avatar TEXT,
    source VARCHAR(50),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE public.tbl_UserOrgLink(
    organisationId INT,
    userId INT
);

ALTER TABLE public.tbl_UserOrgLink
    ADD CONSTRAINT FK_UserOrgLink_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES public.tbl_Organisation (id)
        ON DELETE CASCADE;

ALTER TABLE public.tbl_UserOrgLink
    ADD CONSTRAINT FK_UserOrgLink_UserId 
        FOREIGN KEY (userId)
        REFERENCES public.tbl_User (id)
        ON DELETE CASCADE;

CREATE TABLE public.tbl_ApiClient(
    clientId UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    userId INT,
    secretKey UUID,
    appName VARCHAR(50) NOT NULL,
    description VARCHAR(500),
    validTill TIMESTAMP,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE public.tbl_ApiClient
    ADD CONSTRAINT FK_ApiClient_UserId
    FOREIGN KEY (userId)
    REFERENCES public.tbl_User (id)
    ON DELETE CASCADE;

CREATE TABLE public.tbl_TimeEntry(
    id BIGSERIAL PRIMARY KEY,
    comments VARCHAR(200),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    value NUMERIC NOT NULL,
    valueType varchar(20) NOT NULL,
    userId INT,
    organisationId INT
);

CREATE INDEX IX_timeentry_created ON public.tbl_TimeEntry (created);
CREATE INDEX IX_timeentry_updated ON public.tbl_TimeEntry (updated);

ALTER TABLE public.tbl_TimeEntry
    ADD CONSTRAINT FK_TimeEntry_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES public.tbl_Organisation (id);

ALTER TABLE public.tbl_TimeEntry
    ADD CONSTRAINT FK_TimeEntry_UserId 
        FOREIGN KEY (userId)
        REFERENCES public.tbl_User (id);

CREATE TABLE public.tbl_RepoItem(
    id BIGSERIAL PRIMARY KEY,
    itemIdSource VARCHAR(200) NOT NULL,
    itemType VARCHAR(50) NOT NULL,
    source VARCHAR(50) NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    description VARCHAR(2000),
    timeEntryId BIGINT
);

CREATE INDEX IX_repoitem_source ON tbl_RepoItem (source);

ALTER TABLE public.tbl_RepoItem
    ADD CONSTRAINT FK_RepoItem_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES public.tbl_TimeEntry (id)
        ON DELETE CASCADE;

CREATE TABLE public.tbl_Tag(
    id SERIAL PRIMARY KEY,
    name varchar(50)
);

CREATE TABLE public.tbl_TimeEntryTagLink(
    tagId INT,
    timeEntryId BIGINT
);

ALTER TABLE public.tbl_TimeEntryTagLink
    ADD CONSTRAINT FK_TimeEntryTagLink_TagId 
        FOREIGN KEY (tagId)
        REFERENCES public.tbl_Tag (id)
        ON DELETE CASCADE;

ALTER TABLE public.tbl_TimeEntryTagLink
    ADD CONSTRAINT FK_TimeEntryTagLink_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES public.tbl_TimeEntry (id)
        ON DELETE CASCADE;