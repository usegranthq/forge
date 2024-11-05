-- Modify "oidc_clients" table
ALTER TABLE "oidc_clients" ADD COLUMN "audience" character varying NOT NULL;
-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "provider" character varying NOT NULL;
