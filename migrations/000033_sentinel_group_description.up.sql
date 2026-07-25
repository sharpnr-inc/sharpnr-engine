-- ============================================================
-- Add an optional description to sentinel_engine.groups
--
-- Groups now carry a human-facing description (mirrors roles.description), so a
-- group like the default "Stakeholders" group seeded at provisioning can explain
-- its purpose in the UI. Nullable: existing groups need no backfill.
-- ============================================================

ALTER TABLE sentinel_engine.groups
    ADD COLUMN description TEXT;
