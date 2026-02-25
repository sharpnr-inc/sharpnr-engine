# Organization Signup (`org_signup`)

## Overview

This package manages the creation of new organizations and the assignment of their default Super Admin.

## Features

-   **Public Routes**: Exposes public API endpoints for organization sign-ups.
-   **Service Integration**: Works with the mailer engine, organization service, and employee service.

## Workflow

1.  **Initial Step**: A user is redirected to the orchestrator page.
2.  **Verification**: The system checks if the organization has been set up.
3.  **Next Step**: The user is guided to the Setup Wizard to complete the process.
