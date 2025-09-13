-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL,
  "name" text NULL,
  "email" text NULL,
  PRIMARY KEY ("id")
);
-- Create "type_groups" table
CREATE TABLE "type_groups" (
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_by" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_type_groups_created_by" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_type_groups_updated_by" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "accounts" table
CREATE TABLE "accounts" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_by" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "balance" numeric(10,2) NOT NULL DEFAULT 0,
  "initial_balance" numeric(10,2) NOT NULL DEFAULT 0,
  "type_group_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_accounts_created_by" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_accounts_type_group" FOREIGN KEY ("type_group_id") REFERENCES "type_groups" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_accounts_updated_by" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_accounts_user_id" to table: "accounts"
CREATE INDEX "idx_accounts_user_id" ON "accounts" ("user_id");
-- Create "types" table
CREATE TABLE "types" (
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_by" uuid NOT NULL,
  "type_group_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_types_created_by" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_types_type_group" FOREIGN KEY ("type_group_id") REFERENCES "type_groups" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_types_updated_by" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "categories" table
CREATE TABLE "categories" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_by" uuid NOT NULL,
  "type_id" bigint NOT NULL,
  "parent_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_categories_children" FOREIGN KEY ("parent_id") REFERENCES "categories" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_categories_created_by" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_categories_type" FOREIGN KEY ("type_id") REFERENCES "types" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_categories_updated_by" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "transactions" table
CREATE TABLE "transactions" (
  "workspace_id" uuid NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_by" uuid NOT NULL,
  "amount" numeric(10,2) NOT NULL DEFAULT 0,
  "transaction_date" timestamptz NOT NULL,
  "category_id" bigint NOT NULL,
  "source_id" bigint NOT NULL,
  "destination_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_transactions_category" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_transactions_created_by" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_transactions_destination" FOREIGN KEY ("destination_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_transactions_source" FOREIGN KEY ("source_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_transactions_updated_by" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_transactions_workspace_id" to table: "transactions"
CREATE INDEX "idx_transactions_workspace_id" ON "transactions" ("workspace_id");
