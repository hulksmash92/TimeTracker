-- Create the timetracker DB and switch to that db onto the public schema
CREATE DATABASE IF NOT EXISTS timetracker;
SET SEARCH_PATH TO timetracker;
SET SEARCH_PATH TO public;

-- Add the UUID extension to our DB
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.tbl_User(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(400),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    githubUserId VARCHAR(100),
    avatar TEXT
);

-- Add an index to the user table on the githubUserId column as 
-- we will be commonly querying on this column, 
-- for example checking if a user exists in the db after 
-- they've authenticated using GitHub OAuth
CREATE INDEX IF NOT EXISTS IX_user_githubUserId ON public.tbl_User (githubuserid);

CREATE TABLE IF NOT EXISTS public.tbl_Organisation(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description varchar(500),
    avatar TEXT,
    source VARCHAR(50),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create a link table for linking User <-> Organisation
-- Note: users can be members of multiple organisations, 
--       so they may multiple org links in this table
CREATE TABLE IF NOT EXISTS public.tbl_UserOrgLink(
    organisationId INT,
    userId INT,
    CONSTRAINT FK_UserOrgLink_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES public.tbl_Organisation (id)
        ON DELETE CASCADE,
    CONSTRAINT FK_UserOrgLink_UserId 
        FOREIGN KEY (userId)
        REFERENCES public.tbl_User (id)
        ON DELETE CASCADE
);

-- This table will likely be used in later phases of the application to
-- allow user's to create credentials for daemon applications that may 
-- extract user data using the API.
CREATE TABLE IF NOT EXISTS public.tbl_ApiClient(
    clientId UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    userId INT,
    secretKey VARCHAR(200),
    appName VARCHAR(50) NOT NULL,
    description VARCHAR(500),
    validTill TIMESTAMP,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT FK_ApiClient_UserId 
        FOREIGN KEY (userId)
        REFERENCES public.tbl_User (id)
        ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS public.tbl_TimeEntry(
    id BIGSERIAL PRIMARY KEY,
    comments VARCHAR(200),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    value NUMERIC NOT NULL,
    valueType varchar(20) NOT NULL,
    userId INT,
    organisationId INT,
    CONSTRAINT FK_TimeEntry_OrganisationId 
        FOREIGN KEY (organisationId)
        REFERENCES public.tbl_Organisation (id),
    CONSTRAINT FK_TimeEntry_UserId 
        FOREIGN KEY (userId)
        REFERENCES public.tbl_User (id)
);

CREATE INDEX IF NOT EXISTS IX_timeentry_created ON public.tbl_TimeEntry (created);
CREATE INDEX IF NOT EXISTS IX_timeentry_updated ON public.tbl_TimeEntry (updated);

CREATE IF NOT EXISTS TABLE public.tbl_RepoItem(
    id BIGSERIAL PRIMARY KEY,
    itemIdSource VARCHAR(200) NOT NULL,
    itemType VARCHAR(50) NOT NULL,
    source VARCHAR(50) NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    repoName VARCHAR(200) NOT NULL,
    description VARCHAR(2000),
    timeEntryId BIGINT,
    CONSTRAINT FK_RepoItem_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES public.tbl_TimeEntry (id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS IX_repoitem_source ON tbl_RepoItem (source);

CREATE TABLE IF NOT EXISTS public.tbl_Tag(
    id SERIAL PRIMARY KEY,
    name varchar(50),
    userId INT,
    CONSTRAINT FK_Tag_UserId
        FOREIGN KEY (userId)
        REFERENCES public.tbl_User (id)
);

-- links our tags to the time entries
CREATE TABLE IF NOT EXISTS public.tbl_TimeEntryTagLink(
    tagId INT,
    timeEntryId BIGINT,
    CONSTRAINT FK_TimeEntryTagLink_TagId 
        FOREIGN KEY (tagId)
        REFERENCES public.tbl_Tag (id)
        ON DELETE CASCADE,
    CONSTRAINT FK_TimeEntryTagLink_TimeEntryId 
        FOREIGN KEY (timeEntryId)
        REFERENCES public.tbl_TimeEntry (id)
        ON DELETE CASCADE
);

-- Add some global tags for time entries
INSERT INTO public.tbl_Tag (name)
VALUES ('Feature')
    , ('Testing')
    , ('Design')
    , ('Proof of concept')
    , ('Bug Fix')
    , ('Research')

-- Updates the user details
CREATE OR REPLACE PROCEDURE sp_user_update (
    user_id BIGINT,
    new_name VARCHAR(200),
    new_email VARCHAR(400)
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF TRIM(BOTH FROM COALESCE(new_name, '')) = '' THEN
        new_name = NULL;
    END IF;
    IF TRIM(BOTH FROM COALESCE(new_name, '')) = '' THEN
        new_email = NULL;
    END IF;

    UPDATE tbl_user
    SET updated = NOW(),
        name = coalesce(new_name, githubuserid),
        email = new_email
    WHERE id = user_id;

END;$$;