CREATE TABLE IF NOT EXISTS "authors" (
	"id" INTEGER NOT NULL,
	"nickname" VARCHAR NOT NULL,
	"avatar_file" VARCHAR,
	"email" VARCHAR,
	"github" VARCHAR,
	"twitter" VARCHAR,
	"bilibili" VARCHAR,
	"active" BOOLEAN,
	"create_at" TIMESTAMP,
	PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "categories" (
	"id" INTEGER NOT NULL,
	"parent_id" INTEGER,
	"name" VARCHAR,
	"description" TEXT,
	"create_at" TIMESTAMP,
	PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "post_tags" (
	"id" INTEGER NOT NULL,
	"post_id" INTEGER,
	"tag_id" INTEGER,
	"create_at" TIMESTAMP,
	PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "tags" (
	"id" INTEGER NOT NULL,
	"name" VARCHAR,
	"usage_count" INTEGER,
	"create_at" TIMESTAMP,
	PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "posts" (
	"id" INTEGER NOT NULL,
	"author_id" INTEGER,
	"categories_id" INTEGER,
	"title" VARCHAR,
	"body" TEXT,
	"publish_at" TIMESTAMP,
	"status" INTEGER,
	"view_count" INTEGER,
	"update_at" TIMESTAMP,
	"create_at" TIMESTAMP,
	PRIMARY KEY("id")
);
