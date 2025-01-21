-- apply schema
BEGIN;
\i /docker-entrypoint-initdb.d/schema/001_init_db.sql
COMMIT;

-- run seeds
BEGIN;
\i /docker-entrypoint-initdb.d/seeds/init_db.sql
COMMIT;
