package constants

// --------------------
// Cookie / Token Names
// --------------------
const (
	TOKEN_KEY_REGISTRATION  = "_sharpnr_reg_token_"    // Step 1: initial register
	TOKEN_KEY_VERIFY_OTP    = "_sharpnr_verify_token_" // Step 2: OTP verified
	TOKEN_KEY_AUTHENTICATED = "_sharpnr_auth_token_"   // Step 3: fully registered / logged in
)

// --------------------
// Token Durations (seconds)
// --------------------
const (
	TOKEN_DURATION_SHORT  = 60 * 10       // 10 minutes (example for OTP)
	TOKEN_DURATION_MEDIUM = 3600 * 24     // 1 day
	TOKEN_DURATION_LONG   = 3600 * 24 * 7 // 7 days
)

// --------------------
// Signup Stages
// --------------------
const (
	SIGNUP_STAGE_NONE          = "none"
	SIGNUP_STAGE_REGISTERED    = "registered"
	SIGNUP_STAGE_OTP_VERIFIED  = "otp_verified"
	SIGNUP_STAGE_AUTHENTICATED = "authenticated"
)
