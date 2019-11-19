
DROP TRIGGER amis_updated_at on amis;
ALTER TABLE scans DROP COLUMN ami_id;

DROP TABLE amis;
