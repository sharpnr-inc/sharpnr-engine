-- +migrate Up

-- 1. Add the new column
ALTER TABLE accounting_engine.account
ADD COLUMN accounting_book_id UUID;

-- 2. Backfill accounting_book_id from account_group
UPDATE accounting_engine.account AS a
SET accounting_book_id = ag.accounting_book_id
FROM accounting_engine.account_group AS ag
WHERE a.account_group_id = ag.id;

-- 3. Make column NOT NULL after data is populated
ALTER TABLE accounting_engine.account
ALTER COLUMN accounting_book_id SET NOT NULL;

-- 4. Create index
CREATE INDEX idx_account_accounting_book_id
ON accounting_engine.account(accounting_book_id);