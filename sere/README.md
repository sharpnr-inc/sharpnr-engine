# **Sharpnr Event Relay Engine (SERE)**

## 🧩 Introduction

The **Sharpnr Event Relay Engine (SERE)** is a core package designed to **listen for specific events**, **process them**, and **relay** those events to all registered **subscribers** in real-time.

SERE serves as the **event backbone** for the Sharpnr ecosystem — enabling **reactive, event-driven** communication between microservices and clients. It also provides **WebSocket integration** for seamless live data streaming to connected users or applications.

---

## 🚀 Core Features

- **Event Listening** – Continuously monitors and captures system or custom-defined events.  
- **Event Publishing** – Dispatches captured events to all active subscribers.  
- **WebSocket Integration** – Real-time communication layer for browser and app clients.  
- **Scalable Architecture** – Designed to handle multiple channels and event types efficiently.  
- **Decoupled Design** – Enables modular and independent service communication.  
- **Subscriber Management** – Provides methods for adding, removing, and notifying subscribers.

---

## 🏗️ Architecture Overview

```text
+---------------------------+
|       Event Source        |
| (e.g., DB, API, Service)  |
+------------+--------------+
             |
             v
+---------------------------+
|       SERE Listener       |
|  - Listens for events     |
|  - Parses & validates     |
+------------+--------------+
             |
             v
+---------------------------+
|      Event Dispatcher     |
|  - Routes event to subs   |
|  - Manages delivery       |
+------------+--------------+
             |
             v
+---------------------------+
|   WebSocket Broadcaster   |
|  - Pushes updates live    |
|  - Handles connections    |
+---------------------------+
