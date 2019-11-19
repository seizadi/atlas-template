
DROP TRIGGER vpcs_updated_at on vpcs;
ALTER TABLE ec2s DROP COLUMN vpc_id;

DROP TABLE vpcs;
