package org_signup_services

import "context"

type OrgSignUpServiceInterface interface {

	// ─────────────────────────────
	// STEP 1: INITIATE SIGNUP
	// ─────────────────────────────
	// - Validate input
	// - Check main DB (email / phone / slug)
	// - Create temp org in Redis (TTL)
	// - Generate & send OTP
	// - Return signup_id (tracking key)
	InitiateSignup(
		ctx context.Context)
}
