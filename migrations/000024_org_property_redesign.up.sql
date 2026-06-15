-- ==========================================
-- Remove obsolete columns and tables
-- ==========================================

-- Remove isConfigured column
ALTER TABLE organization_property.organization_property
DROP COLUMN IF EXISTS isConfigured;

-- Add deleted_at column
ALTER TABLE organization_property.organization_property
ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ;

-- ===================================================
-- The following tables were already manually deleted.
-- Keeping these statements commented for migration
-- history purposes.
-- ===================================================

-- DROP TABLE organization_property.organization_property_domain;
-- DROP TABLE organization_property.organization_property_domain_modules;
-- DROP TABLE organization_property.organization_property_domain_infra;