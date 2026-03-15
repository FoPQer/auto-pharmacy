DROP TRIGGER IF EXISTS soft_delete_tag_trigger ON tags;
DROP TRIGGER IF EXISTS update_tags_updated_at_trigger ON tags;

DROP FUNCTION IF EXISTS soft_delete_tag();
DROP FUNCTION IF EXISTS update_tags_updated_at();

DROP INDEX IF EXISTS idx_tags_deleted_at;

DROP TABLE IF EXISTS tags;
