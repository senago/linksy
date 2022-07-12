CREATE UNLOGGED TABLE IF NOT EXISTS url(
  "hash" CHAR(10) PRIMARY KEY,
  "value" text NOT NULL,
  "created_at" timestamp with time zone NOT NULL,
  "expires_at" timestamp with time zone NOT NULL
);

VACUUM ANALYZE;
