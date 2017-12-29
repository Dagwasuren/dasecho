CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "articles" (
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
"id" TEXT PRIMARY KEY,
"title" text NOT NULL,
"content" text NOT NULL,
"author" text NOT NULL,
"uid" integer NOT NULL
);
CREATE TABLE IF NOT EXISTS "todaybests" (
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
"id" TEXT PRIMARY KEY,
"content" TEXT NOT NULL
);
