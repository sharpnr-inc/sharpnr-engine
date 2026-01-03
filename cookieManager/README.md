# cookieManager

cookieManager is an infrastructure-level package that manages **all browser cookies**
used by the Sharpnr Web application.

## Core Principles

- Controllers express **intent**, never mechanics
- Cookie flags, TTLs, and names are **centralized**
- Auth & signup flows are enforced as **state machines**
- Cookies are **never** set ad-hoc in controllers

## What this package manages

- Auth cookies (JWT / session)
- Signup flow cookies (registration → otp → completion)
- CSRF cookies (double-submit pattern)
- Cleanup & transitions

## What this package does NOT do

- JWT signing or verification
- HTTP routing
- CSRF validation middleware
- Frontend logic

## Usage

Initialize once in `main.go` and inject the Manager into controllers.

```go
cookies := cookieManager.New(cookieManager.Config{
    Secure: true,
    Domain: "",
})
