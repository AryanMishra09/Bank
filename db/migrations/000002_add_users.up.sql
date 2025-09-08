CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_change_at" timestamptz NOT NULL DEFAULT '1970-01-01 00:00:00+00',
  "created_at" timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE "account" ADD CONSTRAINT "account_owner_fkey" FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "account" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
