-- +migrate Down

-- 1. Drop index
DROP INDEX IF EXISTS accounting_engine.idx_account_accounting_book_id;

-- 2. Drop column
ALTER TABLE accounting_engine.account
DROP COLUMN IF EXISTS accounting_book_id;