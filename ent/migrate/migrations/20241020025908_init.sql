-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "uid" uuid NOT NULL, "email" character varying NOT NULL, "last_login" timestamptz NULL, "verified_at" timestamptz NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "projects" table
CREATE TABLE "projects" ("id" uuid NOT NULL, "name" character varying NOT NULL, "url_id" character varying NOT NULL, "description" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_projects" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "projects_users_projects" FOREIGN KEY ("user_projects") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "oidc_clients" table
CREATE TABLE "oidc_clients" ("id" uuid NOT NULL, "name" character varying NOT NULL, "client_ref_id" character varying NOT NULL, "client_id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "project_oidc_clients" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "oidc_clients_projects_oidc_clients" FOREIGN KEY ("project_oidc_clients") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "project_domains" table
CREATE TABLE "project_domains" ("id" uuid NOT NULL, "domain" character varying NOT NULL, "verified" boolean NOT NULL DEFAULT false, "verified_at" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "project_domain" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "project_domains_projects_domain" FOREIGN KEY ("project_domain") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "tokens" table
CREATE TABLE "tokens" ("id" uuid NOT NULL, "name" character varying NOT NULL, "token" character varying NOT NULL, "expires_at" timestamptz NULL, "last_used_at" timestamptz NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_tokens" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "tokens_users_tokens" FOREIGN KEY ("user_tokens") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "user_sessions" table
CREATE TABLE "user_sessions" ("id" uuid NOT NULL, "token" character varying NOT NULL, "expires_at" timestamptz NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_user_sessions" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "user_sessions_users_user_sessions" FOREIGN KEY ("user_user_sessions") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "user_verifications" table
CREATE TABLE "user_verifications" ("id" uuid NOT NULL, "attempt_id" uuid NOT NULL, "code" character varying NOT NULL, "attempts" bigint NOT NULL DEFAULT 0, "expires_at" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_user_verifications" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "user_verifications_users_user_verifications" FOREIGN KEY ("user_user_verifications") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
