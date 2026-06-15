-- ============================================================
-- 000024_merge_property_address.down.sql
-- Restore organization_property_address from organization_property
-- ============================================================

-- Step 1: Recreate the address table
CREATE TABLE IF NOT EXISTS organization_property.organization_property_address (
    id                       UUID         PRIMARY KEY,
    organization_id          UUID         NOT NULL,
    organization_property_id UUID         NOT NULL,

    address_line TEXT        NOT NULL,
    city         VARCHAR(100),
    city_code    VARCHAR(20),
    state        VARCHAR(100),
    country      VARCHAR(100),
    pincode      VARCHAR(20),

    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_opa_organization_id
    ON organization_property.organization_property_address (organization_id);

CREATE INDEX IF NOT EXISTS idx_opa_organization_property_id
    ON organization_property.organization_property_address (organization_property_id);

CREATE INDEX IF NOT EXISTS idx_opa_city
    ON organization_property.organization_property_address (city);

CREATE INDEX IF NOT EXISTS idx_opa_city_code
    ON organization_property.organization_property_address (city_code);

CREATE INDEX IF NOT EXISTS idx_opa_state
    ON organization_property.organization_property_address (state);

CREATE INDEX IF NOT EXISTS idx_opa_country
    ON organization_property.organization_property_address (country);

CREATE INDEX IF NOT EXISTS idx_opa_pincode
    ON organization_property.organization_property_address (pincode);

-- Step 2: Restore address data from organization_property
INSERT INTO organization_property.organization_property_address (
    id,
    organization_id,
    organization_property_id,
    address_line,
    city,
    city_code,
    state,
    country,
    pincode,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    organization_id,
    id,
    address_line,
    city,
    city_code,
    state,
    country,
    pincode,
    NOW(),
    NOW()
FROM organization_property.organization_property
WHERE address_line IS NOT NULL;

-- Step 3: Safety check — abort if row counts don't match
DO $$
DECLARE
    property_count INT;
    restored_count INT;
BEGIN
    SELECT COUNT(*) INTO property_count
    FROM organization_property.organization_property
    WHERE address_line IS NOT NULL;

    SELECT COUNT(*) INTO restored_count
    FROM organization_property.organization_property_address;

    IF property_count <> restored_count THEN
        RAISE EXCEPTION 'Data mismatch: % properties with addresses but only % rows restored. Aborting migration.',
            property_count, restored_count;
    END IF;
END $$;

-- Step 4: Drop indexes on organization_property address columns
DROP INDEX IF EXISTS organization_property.idx_op_city;
DROP INDEX IF EXISTS organization_property.idx_op_city_code;
DROP INDEX IF EXISTS organization_property.idx_op_state;
DROP INDEX IF EXISTS organization_property.idx_op_country;
DROP INDEX IF EXISTS organization_property.idx_op_pincode;

-- Step 5: Drop address columns from organization_property
ALTER TABLE organization_property.organization_property
    DROP COLUMN IF EXISTS address_line,
    DROP COLUMN IF EXISTS city,
    DROP COLUMN IF EXISTS city_code,
    DROP COLUMN IF EXISTS state,
    DROP COLUMN IF EXISTS country,
    DROP COLUMN IF EXISTS pincode;