BEGIN;

DROP INDEX "idx_user_username";

DROP TABLE IF EXISTS "users";

COMMIT;