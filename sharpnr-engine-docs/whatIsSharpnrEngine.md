# What is sharpnr-engine?

The `sharpnr-engine` is a robust, multi-tenant CRM (Customer Relationship Management) and ERP (Enterprise Resource Planning) application designed to cater to various business types, each with their unique operational requirements.

## Core Functionality:

At its heart, `sharpnr-engine` provides a comprehensive platform for managing core business processes, including:

*   **Organization Management:** Handling registration, login, and details for different organizations.
*   **Employee Management:** Managing employee data and roles within organizations.
*   **Product and Service Management:** Cataloging and managing products and services offered by organizations.
*   **Transaction Processing:** Facilitating and recording sales and other financial transactions.
*   **Authentication and Authorization:** Securely managing user access and permissions.
*   **Real-time Communication:** Utilizing WebSockets for live updates and interactions.

## Multi-Tenancy and Modular Design:

A key architectural feature of `sharpnr-engine` is its **multi-tenancy** capability. This means a single instance of the application can securely serve multiple distinct organizations (tenants), each operating independently with their own data and configurations.

This is primarily achieved through a **modular design**, evident in the `modules/` directory. The application is structured to support different industry verticals as distinct modules:

*   **`modules/hotel/`:** This module provides specialized CRM/ERP functionalities tailored for the hospitality industry, including booking management, room services, guest management, and housekeeping.
*   **`modules/retail/`:** Designed for retail businesses, this module likely includes features for inventory management, sales processing, supplier management, and point-of-sale (POS) functionalities.
*   **`modules/salon/`:** This module caters to salons, potentially offering features for appointment scheduling, service management, client records, and staff management.

Each module encapsulates its specific business logic, data models, and operational flows, allowing `sharpnr-engine` to be highly adaptable and extensible to new business types.

## Underlying Architecture:

The application leverages a clean architecture pattern, separating concerns into distinct layers:

*   **API Layer (`api/`):** Exposes the application's functionalities via RESTful endpoints.
*   **Controller Layer (`controllers/`):** Processes incoming requests, validates data, and orchestrates calls to the service layer.
*   **Service Layer (`core/services/`, `modules/*/services/`):** Contains the core business logic and rules, independent of the persistence or presentation layers.
*   **Repository Layer (`core/repositories/`, `modules/*/repos/`):** Abstracts data storage details, providing a clean interface for services to interact with the database.
*   **Shared Components (`shared/`, `pkg/`):** Provides common utilities, database connections, middleware, and reusable packages across the application.

This layered approach, combined with the modular and multi-tenant design, makes `sharpnr-engine` a flexible and powerful solution for diverse business needs.