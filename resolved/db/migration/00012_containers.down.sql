
DROP TRIGGER containers_updated_at on containers;
ALTER TABLE scans DROP COLUMN container_id;

DROP TABLE containers;
