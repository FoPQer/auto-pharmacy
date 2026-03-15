CREATE TABLE medicines (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    measurement VARCHAR(50),
    dose FLOAT,
    measurement_dose VARCHAR(20),
    box VARCHAR(255),
    place VARCHAR(255),
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_medicines_deleted_at ON medicines (deleted_at);

CREATE OR REPLACE FUNCTION update_medicines_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_medicines_updated_at
BEFORE UPDATE ON medicines
FOR EACH ROW
EXECUTE FUNCTION update_medicines_updated_at();

CREATE OR REPLACE FUNCTION soft_delete_medicine()
RETURNS TRIGGER AS $$
BEGIN
    NEW.deleted_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER soft_delete_medicines
BEFORE DELETE ON medicines
FOR EACH ROW
EXECUTE FUNCTION soft_delete_medicine();
