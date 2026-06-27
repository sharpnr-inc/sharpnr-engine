package app

// bootstrapSentinelEngine constructs the authorization engine. It wraps the same
// Sentinel repo (which now also implements the engine's read interface) and is
// safe for concurrent use, so a single instance is shared across the app and
// injected into any service that needs to make authorization decisions.
