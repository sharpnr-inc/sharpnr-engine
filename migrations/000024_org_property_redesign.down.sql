-- ==========================================
-- Rollback changes
-- ==========================================

-- Restore isConfigured column
ALTER TABLE organization_property.organization_property
ADD COLUMN IF NOT EXISTS isConfigured BOOLEAN DEFAULT FALSE;

-- Remove deleted_at column
ALTER TABLE organization_property.organization_property
DROP COLUMN IF EXISTS deleted_at;

-- ===================================================
-- Recreate deleted tables (replace with original schema)
-- ===================================================

-- CREATE TABLE organization_property.organization_property_domain (
--     -- original columns here
-- );

-- CREATE TABLE organization_property.organization_property_domain_modules (
--     -- original columns here
-- );

-- CREATE TABLE organization_property.organization_property_domain_infra (
--     -- original columns here
-- );