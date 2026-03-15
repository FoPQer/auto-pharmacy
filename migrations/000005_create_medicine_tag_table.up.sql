CREATE TABLE medicine_tag (
    medicine_id INT NOT NULL REFERENCES medicines(id) ON DELETE CASCADE,
    tag_id INT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (medicine_id, tag_id)
);

CREATE INDEX idx_medicine_tag_medicine_id ON medicine_tag(medicine_id);
CREATE INDEX idx_medicine_tag_tag_id ON medicine_tag(tag_id);