-- Only the two tables this migration created. The accounting_engine schema is
-- NOT dropped: it predates this migration and holds the books, accounts and
-- journals.
DROP TABLE IF EXISTS accounting_engine.accounting_engine_account_settings;
DROP TABLE IF EXISTS accounting_engine.org_accounting_engine_accounts;
