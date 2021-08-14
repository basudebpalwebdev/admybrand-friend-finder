BEGIN;

CREATE TABLE IF NOT EXISTS "users" (
	"id" BIGSERIAL PRIMARY KEY,
	"username" varchar UNIQUE NOT NULL,
	"description" varchar NOT NULL,
	"dob" timestamptz NOT NULL,
	"address" varchar NOT NULL,
	"created_at" timestamptz NOT NULL
);

CREATE UNIQUE INDEX "idx_user_username" ON "users" ("username");

COMMIT;