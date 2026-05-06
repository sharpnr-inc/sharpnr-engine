# Sharpnr Engine

**Multi-tenant SaaS Orchestration & ERP Core**

Sharpnr is a high-performance orchestration engine designed to coordinate complex workflows across distributed modules. It serves as the backbone for hospitality, accounting, and human resource systems, ensuring data integrity and cross-domain communication.

---

## 1. System Architecture

Sharpnr operates as a centralized orchestrator (API Gateway) that manages the lifecycle of requests and coordinates between specialized domain engines.

```text
┌─────────────────────────────────────────────────────┐
│            ORCHESTRATOR (API Gateway)                │
│     Coordinates workflows across domains             │
└──────────────┬──────────────┬────────────────────────┘
               │              │
    ┌──────────┴──────────┐   └──────────────┐
    │  DOMAIN ENGINES     │                  │
    ├─────────────────────┤         ┌────────▼────────┐
    │ modules/hr/         │         │ sharp_services/ │
    │ modules/accounting/ │         │ (Internal/Auth) │
    │ modules/hospitality/│         └─────────────────┘
    └──────────┬──────────┘
               │
    ┌──────────▼──────────┐
    │ DATA PERSISTENCE    │
    │ (PostgreSQL/SQLx)   │
    └─────────────────────┘
```
