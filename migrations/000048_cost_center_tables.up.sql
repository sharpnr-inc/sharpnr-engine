-- --------------------------------------
-- core.org_cost_centers
-- One operating location of an organization — the Bangalore office, the Dubai
-- branch, a warehouse. An organization has as many as it needs, in as many
-- countries as it needs.
--
-- legal_entity_id is nullable so a site can be stood up before its entity is
-- registered, but it must be set before anything here can be invoiced or filed:
-- an office in another country almost always belongs to a different entity than
-- head office, and that is what decides tax treatment.
--
-- timezone is per cost center, not per organization — attendance cutoffs, shift
-- boundaries and payroll periods are all local to the site.
-- --------------------------------------
CREATE TABLE core.org_cost_centers (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    legal_entity_id UUID,

    code            VARCHAR(50) NOT NULL,
    name            VARCHAR(150) NOT NULL,
    description     TEXT,

    type            VARCHAR(30),
    status          VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',

    address_line    TEXT,
    city            VARCHAR(100),
    state           VARCHAR(100),
    country         VARCHAR(100),
    pincode         VARCHAR(20),

    timezone        VARCHAR(50),

    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

-- The tenant predicate every read carries.
CREATE INDEX idx_org_cost_centers_organization_id ON core.org_cost_centers(organization_id);

-- "which cost centers file under this entity" — the statutory reporting roll-up.
CREATE INDEX idx_org_cost_centers_legal_entity_id ON core.org_cost_centers(legal_entity_id);

CREATE INDEX idx_org_cost_centers_code       ON core.org_cost_centers(code);
CREATE INDEX idx_org_cost_centers_type       ON core.org_cost_centers(type);
CREATE INDEX idx_org_cost_centers_status     ON core.org_cost_centers(status);
CREATE INDEX idx_org_cost_centers_deleted_at ON core.org_cost_centers(deleted_at);

-- Location filters for the org's cost-center list view.
CREATE INDEX idx_org_cost_centers_city    ON core.org_cost_centers(city);
CREATE INDEX idx_org_cost_centers_state   ON core.org_cost_centers(state);
CREATE INDEX idx_org_cost_centers_country ON core.org_cost_centers(country);
CREATE INDEX idx_org_cost_centers_pincode ON core.org_cost_centers(pincode);

-- (org, code) is unique among live rows; a soft-deleted row must not block
-- re-using the same code, hence the partial predicate.
CREATE UNIQUE INDEX uq_org_cost_centers_org_code
    ON core.org_cost_centers(organization_id, code)
    WHERE deleted_at IS NULL;

-- --------------------------------------
-- core.cost_center_departments
-- A department belonging to one cost center — the Bangalore office's Finance
-- team. A cost center has as many as it needs.
--
-- These are the cost center's OWN departments, not placements of an org-level
-- department. The two department tracks are independent by design: an org that
-- manages cost centers defines its departments here, and an org that does not
-- uses core.departments, where departments hang off the organization directly.
--
-- organization_id is denormalised from the cost center so every tenant-scoped
-- read of this table is a single indexed predicate rather than a join.
-- --------------------------------------
CREATE TABLE core.cost_center_departments (
    id                 UUID PRIMARY KEY,
    organization_id    UUID NOT NULL,
    org_cost_center_id UUID NOT NULL,

    code               VARCHAR(50) NOT NULL,
    name               VARCHAR(150) NOT NULL,
    description        TEXT,

    head_employee_id   UUID,

    status             VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',

    created_at         TIMESTAMPTZ DEFAULT now(),
    updated_at         TIMESTAMPTZ DEFAULT now(),
    deleted_at         TIMESTAMPTZ
);

CREATE INDEX idx_cost_center_departments_organization_id    ON core.cost_center_departments(organization_id);

-- "list the departments of this cost center" — the primary read.
CREATE INDEX idx_cost_center_departments_org_cost_center_id ON core.cost_center_departments(org_cost_center_id);

CREATE INDEX idx_cost_center_departments_code             ON core.cost_center_departments(code);
CREATE INDEX idx_cost_center_departments_head_employee_id ON core.cost_center_departments(head_employee_id);
CREATE INDEX idx_cost_center_departments_status           ON core.cost_center_departments(status);
CREATE INDEX idx_cost_center_departments_deleted_at       ON core.cost_center_departments(deleted_at);

-- (cost center, code) is unique among live rows; a soft-deleted row must not
-- block re-using the same code, hence the partial predicate.
CREATE UNIQUE INDEX uq_cost_center_departments_cost_center_code
    ON core.cost_center_departments(org_cost_center_id, code)
    WHERE deleted_at IS NULL;
