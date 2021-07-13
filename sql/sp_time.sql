-- View for getting time entries with basic user and org details
CREATE OR REPLACE VIEW vw_time_entries
AS (
    SELECT t.id, t.created, t.updated, t.value, t.valueType,
        coalesce(t.comments, '') AS comments,
        coalesce(t.userId, 0) AS userId,
        coalesce(u.name, '') AS username,
        coalesce(u.avatar, '') AS userAvatar,
        coalesce(t.organisationId, 0) AS organisationId,
        coalesce(o.name, '') AS organisation,
        coalesce(o.avatar, '') AS organisationAvatar
    FROM tbl_timeentry AS t
    LEFT JOIN tbl_user AS u ON u.id = t.userId
    LEFT JOIN tbl_organisation AS o ON o.id = t.organisationId
);

-- View for getting tags that are linked to time entries
CREATE OR REPLACE VIEW vw_time_entry_tags
AS (
    SELECT t.id, t.name, t.userId, ttl.timeentryid
    FROM tbl_timeentrytaglink AS ttl
    INNER JOIN tbl_tag AS t ON t.id = ttl.tagid
);

-- View for getting repo item with null column values filled in
CREATE OR REPLACE VIEW vw_repo_items
AS (
    SELECT r.id, r.created, r.updated, r.itemidsource, r.itemtype, r.source, r.reponame,
        coalesce(r.description, '') AS description, r.timeentryid
    FROM tbl_repoitem AS r
);

-- Inserts a new time entry and outputs the new time entry ID
CREATE OR REPLACE PROCEDURE sp_time_insert (
    arg_user_id BIGINT,
    arg_org_id BIGINT,
    arg_comments VARCHAR(200),
    arg_value NUMERIC,
    arg_valueType VARCHAR(20),
    timeentryid INOUT BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    IF arg_user_id = 0 THEN
        arg_user_id := NULL;
    END IF;
    IF arg_org_id = 0 THEN
        arg_org_id := NULL;
    END IF;
    IF arg_comments = '' THEN
        arg_comments := NULL;
    END IF;

    INSERT INTO tbl_timeentry (
        userId,
        organisationId,
        comments,
        value,
        valueType
    )
    VALUES (
        arg_user_id,
        arg_org_id,
        arg_comments,
        arg_value,
        arg_valueType
    )
    RETURNING id INTO timeentryid;
END;$$;

-- Deletes a time entry
CREATE OR REPLACE PROCEDURE sp_time_delete (entry_id BIGINT)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM tbl_timeentry WHERE id = entry_id;
END;$$;

-- Creates a link between a time entry and a tag, or creates a new tag if one isn't found
-- TODO: Add userId param for user created tags
CREATE OR REPLACE PROCEDURE sp_time_tags_insert (
    time_entry_id BIGINT,
    tag_name VARCHAR(50),
    tag_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    SELECT id INTO tag_id FROM tbl_tag WHERE name = tag_name;

    IF tag_id = 0 OR tag_id IS NULL THEN
        INSERT INTO tbl_tag (name)
        VALUES (tag_name)
        RETURNING id INTO tag_id;
    END IF;

    INSERT INTO tbl_timeentrytaglink (tagid, timeentryid)
    VALUES (tag_id, time_entry_id);
END;$$;

-- Deletes a time entry tag links
CREATE OR REPLACE PROCEDURE sp_time_tags_delete (time_entry_id BIGINT, tag_id BIGINT)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM tbl_timeentrytaglink WHERE tagid = tag_id AND timeentryid = time_entry_id;
END;$$;

-- deletes a repo item with the given row ID
CREATE OR REPLACE PROCEDURE sp_time_repoitem_delete (item_id BIGINT)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM tbl_repoitem WHERE id = item_id;
END;$$;

-- Creates a new repo item entry in the DB
CREATE OR REPLACE PROCEDURE sp_time_repoitem_insert (
    time_entry_id BIGINT,
    item_created timestamp,
    item_id_source varchar(200),
    item_type varchar(50),
    item_source varchar(50),
    repo_name varchar(200),
    item_desc varchar(2000)
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO tbl_repoitem (
        itemidsource,
        itemtype,
        source,
        created,
        description,
        timeentryid,
        reponame
    )
    VALUES (
        item_id_source,
        item_type,
        item_source,
        item_created,
        item_desc,
        time_entry_id,
        repo_name
    );
END;$$;
