-- ============================================================
-- Store the Cedar JSON form of a policy alongside its text.
--
-- The frontend authors policies with @cedar-policy/cedar-wasm, which can produce
-- both the human-readable text and the structured JSON form. We persist the JSON
-- so the builder can reload/edit a policy without re-parsing. The engine still
-- compiles only sentinel_policies.policy (the text); policy_json is FE storage.
-- Nullable: text-only policies (e.g. the seeded owner policy) have no JSON.
-- ============================================================

BEGIN;

ALTER TABLE sentinel_engine.sentinel_policies
    ADD COLUMN policy_json JSONB;

COMMIT;
