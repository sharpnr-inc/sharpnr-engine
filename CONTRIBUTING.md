# Contributing Guidelines

Thank you for considering contributing to this project! 🚀

We follow a layered architecture designed for **multi-tenant SaaS** supporting various organization types like **Salon**, **Retail**, and **Hotel**. This guide will help you understand how to contribute effectively.

---

## ✅ How the App is Structured

We use a **modular, layered design** built around:

- Models (GORM)
- DTOs (Validation / transport layer)
- Mappers (conversion between DTO ↔ model)
- Repositories (database access)
- Services (business logic)
- Controllers (Gin handlers)
- Routers (routes grouped per module or API version)
- Modules (org-type-specific logic: Salon, Hotel, Retail)

---

## 📦 Folder Structure (Simplified)


---

## 🧠 Application Architecture

```mermaid
flowchart TD
  A[Client Request] --> B[Router (Gin)]
  B --> C[Controller]
  C --> D[Service Layer]
  D --> E[Repository Layer]
  E --> F[PostgreSQL DB via GORM]
  D --> G[Mapper (DTO ↔ Model)]
