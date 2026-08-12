DROP TABLE IF EXISTS hrms.org_hrms_preference;

-- No CASCADE: if a later HRMS table is still present the drop fails loudly
-- rather than silently destroying it.
DROP SCHEMA IF EXISTS hrms;
