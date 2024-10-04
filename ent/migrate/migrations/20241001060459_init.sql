-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "email" character varying NOT NULL, "password" character varying NOT NULL, "last_login" timestamptz NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "projects" table
CREATE TABLE "projects" ("id" uuid NOT NULL, "name" character varying NOT NULL, "description" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_projects" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "projects_users_projects" FOREIGN KEY ("user_projects") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "oidc_clients" table
CREATE TABLE "oidc_clients" ("id" uuid NOT NULL, "name" character varying NOT NULL, "client_id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "project_oidc_clients" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "oidc_clients_projects_oidc_clients" FOREIGN KEY ("project_oidc_clients") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "user_sessions" table
CREATE TABLE "user_sessions" ("id" uuid NOT NULL, "token" character varying NOT NULL, "expires_at" timestamptz NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_user_sessions" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "user_sessions_users_user_sessions" FOREIGN KEY ("user_user_sessions") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
