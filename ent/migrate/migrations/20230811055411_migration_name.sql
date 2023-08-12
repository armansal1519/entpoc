-- Modify "users" table
ALTER TABLE "users" DROP COLUMN "table_name43", DROP COLUMN "table_name50", ADD COLUMN "altered_table_name43" character varying NOT NULL, ADD COLUMN "altered_table_name50" character varying NOT NULL;
