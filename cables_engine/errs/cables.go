// Package cablesEngineErrs defines the sentinel errors returned by the cables
// module. Keeping them in one place lets the service signal a precise failure
// and the controllers map it to a status code with errors.Is, rather than
// matching on message text.
package cablesEngineErrs

import "errors"

var (
	ErrCablesServiceNotInitialized = errors.New("cables service is not started, please initialize it first before using it")
	ErrNoChannelFound              = errors.New("no channel found for the org")
	ErrCablesAccountIsDisabled     = errors.New("cables account is disabled for the org, please contact support")

	// ErrCablesServiceAlreadyStarted answers a second start on an account that
	// is already live. Starting a switched-off account is not this error — it
	// re-activates the existing one.
	ErrCablesServiceAlreadyStarted = errors.New("cables service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets a cables account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the cables service")

	/* -------------------------------- channels ------------------------------- */

	// ErrChannelNameRequired covers both an empty name and one that carries no
	// character a slug can be built from ("###").
	ErrChannelNameRequired = errors.New("channel name is required")

	// ErrChannelAlreadyExists is a name that resolves to a slug another live
	// channel in the same cables account already holds. The partial unique
	// index enforces it underneath as the race backstop.
	ErrChannelAlreadyExists = errors.New("a channel with this name already exists")

	ErrInvalidChannelType       = errors.New("channel type must be one of TEXT, AUDIO or VIDEO")
	ErrInvalidChannelVisibility = errors.New("channel visibility must be either PUBLIC or PRIVATE")

	/* --------------------------- members & messaging -------------------------- */

	// ErrMemberNotFound means the employee has no cables identity yet. Normally
	// impossible after a roster sync, so it points at an employee created since
	// the last one rather than at a bad request.
	ErrMemberNotFound = errors.New("no cables member found for this employee")

	ErrConversationNotFound = errors.New("conversation not found")

	// ErrNotAConversationMember is the authorization answer for every read,
	// send and socket subscribe. It deliberately does not distinguish "does not
	// exist" from "not yours" — that difference is how a caller enumerates
	// conversations it cannot see.
	ErrNotAConversationMember = errors.New("you are not a member of this conversation")

	ErrMessageNotFound = errors.New("message not found")
	ErrMessageEmpty    = errors.New("a message needs a body")
	ErrMessageTooLong  = errors.New("message exceeds the maximum length")

	// ErrChannelNotJoinable answers a join on a private channel. Private means
	// somebody inside has to add you — there is no self-service path in.
	ErrChannelNotJoinable = errors.New("this channel is private, ask a member to add you")

	ErrChannelArchived = errors.New("this channel is archived and read-only")

	// ErrCannotDirectMessageSelf guards the DM upsert: a self-DM would produce
	// a direct_key of one id and a conversation with a single participant.
	ErrCannotDirectMessageSelf = errors.New("you cannot open a direct conversation with yourself")
)
