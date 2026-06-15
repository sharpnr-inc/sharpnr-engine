-- ============================================================
-- 000024_merge_property_address.up.sql
-- Merge organization_property_address into organization_property
-- ============================================================

-- Step 1: Add address columns to organization_property
ALTER TABLE organization_property.organization_property
    ADD COLUMN IF NOT EXISTS address_line TEXT,
    ADD COLUMN IF NOT EXISTS city         VARCHAR(100),
    ADD COLUMN IF NOT EXISTS city_code    VARCHAR(20),
    ADD COLUMN IF NOT EXISTS state        VARCHAR(100),
    ADD COLUMN IF NOT EXISTS country      VARCHAR(100),
    ADD COLUMN IF NOT EXISTS pincode      VARCHAR(20);

-- Step 2: Copy existing address data into the main table
UPDATE organization_property.organization_property op
SET
    address_line = opa.address_line,
    city         = opa.city,
    city_code    = opa.city_code,
    state        = opa.state,
    country      = opa.country,
    pincode      = opa.pincode
FROM organization_property.organization_property_address opa
WHERE opa.organization_property_id = op.id;

-- Step 3: Safety check — abort if row counts don't match
DO $$
DECLARE
    address_count INT;
    copied_count  INT;
BEGIN
    SELECT COUNT(*) INTO address_count
    FROM organization_property.organization_property_address;

    SELECT COUNT(*) INTO copied_count
    FROM organization_property.organization_property
    WHERE address_line IS NOT NULL;

    IF address_count <> copied_count THEN
        RAISE EXCEPTION 'Data mismatch: % addresses exist but only % were copied. Aborting migration.',
            address_count, copied_count;
    END IF;
END $$;

-- Step 4: Add indexes to match original address table
CREATE INDEX IF NOT EXISTS idx_op_city
    ON organization_property.organization_property (city);

CREATE INDEX IF NOT EXISTS idx_op_city_code
    ON organization_property.organization_property (city_code);

CREATE INDEX IF NOT EXISTS idx_op_state
    ON organization_property.organization_property (state);

CREATE INDEX IF NOT EXISTS idx_op_country
    ON organization_property.organization_property (country);

CREATE INDEX IF NOT EXISTS idx_op_pincode
    ON organization_property.organization_property (pincode);

-- Step 5: Drop the address table (only reached if safety check passed)
DROP TABLE IF EXISTS organization_property.organization_property_address;