CREATE TABLE supplies (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    medicine_id BIGINT REFERENCES medicines(id) ON DELETE CASCADE,
    quantity INTEGER,
    expired_at DATE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE OR REPLACE FUNCTION update_supplies_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_supplies_updated_at_trigger
BEFORE UPDATE ON supplies
FOR EACH ROW
EXECUTE FUNCTION update_supplies_updated_at();

CREATE OR REPLACE FUNCTION soft_delete_supply()
RETURNS TRIGGER AS $$
BEGIN
    NEW.deleted_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER soft_delete_supply_trigger
BEFORE DELETE ON supplies
FOR EACH ROW
EXECUTE FUNCTION soft_delete_supply();

CREATE INDEX idx_supplies_deleted_at ON supplies(deleted_at);
CREATE INDEX idx_supplies_medicine_id ON supplies(medicine_id);