DROP TRIGGER IF EXISTS soft_delete_medicines ON medicines;
DROP TRIGGER IF EXISTS update_medicines_updated_at ON medicines;

DROP FUNCTION IF EXISTS soft_delete_medicine();
DROP FUNCTION IF EXISTS update_medicines_updated_at();

DROP INDEX IF EXISTS idx_medicines_deleted_at;

DROP TABLE IF EXISTS medicines;
