package cablesEngineService

import (
	"context"
	"log/slog"

	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	ws_engine "sharpnr-engine/ws_engine"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                             Topic authorization                            */
/* -------------------------------------------------------------------------- */

// ConversationTopicAuthorizer answers the websocket engine's question for the
// "conv:" namespace: may this identity listen to this conversation?
//
// The engine has no idea what a conversation is — this is the only place the
// rule lives, and it is the same membership check the HTTP reads run. A socket
// is not a second, weaker way into the data.
type ConversationTopicAuthorizer struct {
	service *cablesService
}

func (a *ConversationTopicAuthorizer) CanSubscribe(ctx context.Context, id ws_engine.Identity, topic string) error {
	conversationID, err := parseTopicUUID(topic)
	if err != nil {
		return err
	}

	_, _, err = a.service.authorizeConversation(ctx, id.OrgID, id.UserID, conversationID)
	return err
}

/* ---------------------------------- <***> --------------------------------- */

// MemberTopicAuthorizer guards the "member:" namespace — a member's private
// channel for DM invites and mention pings.
//
// The rule is simply "this must be you": the topic id is a cables member id,
// and it has to resolve to the employee behind the socket.
type MemberTopicAuthorizer struct {
	service *cablesService
}

func (a *MemberTopicAuthorizer) CanSubscribe(ctx context.Context, id ws_engine.Identity, topic string) error {
	memberID, err := parseTopicUUID(topic)
	if err != nil {
		return err
	}

	me, err := a.service.resolveMember(ctx, id.OrgID, id.UserID)
	if err != nil {
		return err
	}

	if me.ID != memberID {
		return cablesEngineErrsNotYourTopic
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// RegisterRealtimeTopics claims cables' websocket namespaces on the engine.
//
// Called once at bootstrap. Until it runs, the engine denies every "conv:" and
// "member:" subscribe — namespaces are opt-in, so a cables build without this
// call fails closed rather than exposing the topics.
func RegisterRealtimeTopics(engine *ws_engine.Engine, service CablesEngineServiceInterface) {
	concrete, ok := service.(*cablesService)
	if !ok || engine == nil {
		return
	}

	engine.RegisterAuthorizer(cablesEngineConstants.TopicNamespaceConversation, &ConversationTopicAuthorizer{service: concrete})
	engine.RegisterAuthorizer(cablesEngineConstants.TopicNamespaceMember, &MemberTopicAuthorizer{service: concrete})
}

/* -------------------------------------------------------------------------- */
/*                                 Publishing                                 */
/* -------------------------------------------------------------------------- */

// publishToConversation fans an event out to everyone watching a conversation.
//
// Always called AFTER the transaction it announces has committed — a client
// that hears about a message before it is readable will fetch nothing and
// render a hole.
func (c *cablesService) publishToConversation(conversationID uuid.UUID, event string, payload any) {
	if c.wsEngine == nil {
		return
	}

	c.wsEngine.Publish(ws_engine.Event{
		Topic: cablesEngineConstants.TopicNamespaceConversation + ":" + conversationID.String(),
		Name:  event,
		Data:  payload,
	})
}

// publishToMember reaches one person wherever they are connected, for things
// addressed to them rather than to a room.
func (c *cablesService) publishToMember(cablesMemberID uuid.UUID, event string, payload any) {
	if c.wsEngine == nil {
		return
	}

	c.wsEngine.Publish(ws_engine.Event{
		Topic: cablesEngineConstants.TopicNamespaceMember + ":" + cablesMemberID.String(),
		Name:  event,
		Data:  payload,
	})
}

// logRealtimeFailure records a best-effort step that failed without failing the
// request around it. Used for the bookkeeping that is allowed to lose: a badge
// that self-corrects, a notification nobody is waiting on.
func (c *cablesService) logRealtimeFailure(what string, err error) {
	slog.Warn("cables best-effort step failed", "step", what, "error", err)
}
