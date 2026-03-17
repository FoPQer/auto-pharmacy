CREATE TABLE tags (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE OR REPLACE FUNCTION update_tags_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_tags_updated_at_trigger
BEFORE UPDATE ON tags
FOR EACH ROW
EXECUTE FUNCTION update_tags_updated_at();

CREATE OR REPLACE FUNCTION soft_delete_tag()
RETURNS TRIGGER AS $$
BEGIN
    NEW.deleted_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER soft_delete_tag_trigger
BEFORE DELETE ON tags
FOR EACH ROW
EXECUTE FUNCTION soft_delete_tag();

CREATE INDEX idx_tags_deleted_at ON tags(deleted_at);