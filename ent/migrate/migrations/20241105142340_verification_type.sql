-- Create "verifications" table
CREATE TABLE "verifications" ("id" uuid NOT NULL, "attempt_id" uuid NOT NULL, "type" character varying NOT NULL, "code" character varying NOT NULL, "attempts" bigint NOT NULL DEFAULT 0, "expires_at" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_verifications" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "verifications_users_verifications" FOREIGN KEY ("user_verifications") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Drop "user_verifications" table
DROP TABLE "user_verifications";
