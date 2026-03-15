DROP TRIGGER IF EXISTS soft_delete_supply_trigger ON supplies;
DROP TRIGGER IF EXISTS update_supplies_updated_at_trigger ON supplies;

DROP FUNCTION IF EXISTS soft_delete_supply();
DROP FUNCTION IF EXISTS update_supplies_updated_at();

DROP INDEX IF EXISTS idx_supplies_medicine_id;
DROP INDEX IF EXISTS idx_supplies_deleted_at;

DROP TABLE IF EXISTS supplies;
