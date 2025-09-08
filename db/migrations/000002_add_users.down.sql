-- Drop the UNIQUE constraint on "account"
ALTER TABLE "account" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- Drop the FOREIGN KEY on "account"
ALTER TABLE "account" DROP CONSTRAINT IF EXISTS account_owner_fkey;

-- Drop the "users" table
DROP TABLE IF EXISTS "users";
