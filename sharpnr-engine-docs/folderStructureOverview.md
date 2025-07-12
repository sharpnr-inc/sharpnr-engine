# Project Folder Structure Overview

This document provides an overview of the `sharpnr-engine` project's folder structure, explaining the purpose and responsibilities of each main directory.

```
sharpnr-engine/
├── api/
├── cmd/
├── config/
├── controllers/
├── core/
├── docs/
├── modules/
├── pkg/
├── shared/
├── terminals/
└── websocket/
```

---

## Directory Breakdown:

### `api/`
*   **Purpose:** Defines the API routes and handles the routing of incoming HTTP requests to the appropriate controllers. It acts as the entry point for external communication with the application.
*   **Contents:**
    *   `router.go`: Main router setup.
    *   `routes/v1/api.go`: Versioned API route definitions.
    *   `v1/`: Contains specific route definitions for API version 1, including core and module-specific routes.

### `cmd/`
*   **Purpose:** Contains the main entry points for executable applications within the project.
*   **Contents:**
    *   `server/`: Holds the main server application (`main.go`), WebSocket handling, and worker processes.

### `config/`
*   **Purpose:** Manages application-wide configuration settings.
*   **Contents:**
    *   `config.go`: General application configuration.
    *   `module_config.go`: Configuration specific to different modules.
    *   `terminal_config.go`: Configuration for terminal types.
    *   `websocket_config.go`: Configuration for WebSocket connections.

### `controllers/`
*   **Purpose:** Handles incoming HTTP requests, validates input, delegates business logic to services, and formats responses. Controllers are the interface between the HTTP layer and the application's business logic.
*   **Contents:**
    *   `employeeControllers/`: Handles employee-related API requests.
    *   `organisationControllers/`: Manages organization-related operations (e.g., registration, login, sales, product/service CRUD).
    *   `productControllers/`: Handles product-specific actions.
    *   `transactionControllers/`: Manages transaction creation.

### `core/`
*   **Purpose:** Encapsulates the core business logic, domain models, interfaces, and data transfer objects (DTOs). This is the heart of the application's domain.
*   **Contents:**
    *   `constants/`: Defines application-wide constants (e.g., department types, event types, roles).
    *   `dtos/`: Data Transfer Objects for transferring data between layers (e.g., `employee_dto.go`, `org_Register_dto.go`).
    *   `interfaces/`: Defines interfaces for services, repositories, and domain entities, promoting loose coupling.
    *   `mappers/`: Handles the conversion between DTOs, models, and other data structures.
    *   `models/`: Defines the application's domain models (e.g., `organizationModel.go`, `employeeModel.go`).
    *   `repositories/`: Abstracts data persistence operations, providing an interface for services to interact with the database.
    *   `services/`: Contains the core business logic implementations, orchestrating operations and applying business rules.

### `docs/`
*   **Purpose:** Stores project documentation, API specifications, and setup guides.
*   **Contents:**
    *   `api-specs.md`, `swagger.json`, `swagger.yaml`: API documentation.
    *   `setup.md`: Project setup instructions.
    *   `tenant-strategy.md`: Documentation on multi-tenancy.

### `modules/`
*   **Purpose:** Implements distinct business verticals or modules, each with its own set of domain logic, controllers, services, and repositories. This promotes modularity and scalability for different business needs.
*   **Contents:**
    *   `module_manager.go`: Manages the loading and registration of different modules.
    *   `base/`: Provides base interfaces and structures for modules.
    *   `hotel/`: Contains all components related to the hotel management module (DTOs, mappers, models, controllers, services, repositories, terminals, websockets).
    *   `retail/`: Contains all components related to the retail module (DTOs, mappers, models, services, terminals).
    *   `salon/`: Contains all components related to the salon module (DTOs, mappers, models, services, terminals, websockets).

### `pkg/`
*   **Purpose:** Contains reusable, general-purpose packages that are not specific to the core business domain but provide common functionalities.
*   **Contents:**
    *   `cache/`: Caching mechanisms (e.g., memory, Redis).
    *   `logger/`: Logging utilities.
    *   `pubsub/`: Publish-subscribe messaging system.
    *   `queue/`: Queueing mechanisms.

### `shared/`
*   **Purpose:** Provides common utilities and shared components used across different parts of the application.
*   **Contents:**
    *   `database/`: Database connection and management.
    *   `middlewares/`: HTTP middleware for authentication, user type checking, and tenancy.
    *   `utils/`: General utility functions (e.g., JWT handling, password hashing, unique ID generation).

### `terminals/`
*   **Purpose:** Manages different types of "terminals" or points of interaction within the system, including permissions and registry.
*   **Contents:**
    *   `terminal_manager.go`, `terminal_permission.go`, `terminal_registry.go`: Core terminal management.
    *   `interfaces/`: Defines interfaces for terminal handlers.

### `websocket/`
*   **Purpose:** Handles WebSocket connections and real-time communication within the application.
*   **Contents:**
    *   `broadcast.go`, `client.go`, `connection.go`, `hub.go`, `message.go`: Core WebSocket infrastructure.
    *   `events/`: Defines WebSocket event types.
    *   `handlers/`: Handles specific WebSocket messages and events.
