package cookieManager

import "github.com/gin-gonic/gin"

// Config is provided from main.go
type Config struct {
	Secure bool
	Domain string
}

// Manager is the public API.
type Manager interface {
	Set(ctx *gin.Context, t Type, value string) error
	Clear(ctx *gin.Context, t Type) error

	// 🔥 NEW READ HELPERS
	ReadAuth(ctx *gin.Context) (string, error)
	ReadSignup(ctx *gin.Context) (string, error)
	ReadCSRF(ctx *gin.Context) (string, error)
	Read(ctx *gin.Context, t Type) (string, error)

	// Signup helpers
	SetSignup(ctx *gin.Context, value string) error
	ClearSignup(ctx *gin.Context)
}

type manager struct {
	writer *writer
}

func New(cfg Config) Manager {
	return &manager{
		writer: newWriter(cfg.Secure, cfg.Domain),
	}
}

func (m *manager) Set(ctx *gin.Context, t Type, value string) error {
	def, ok := registry[t]
	if !ok {
		return ErrUnknownCookie
	}
	m.writer.set(ctx, def, value)
	return nil
}

func (m *manager) Clear(ctx *gin.Context, t Type) error {
	def, ok := registry[t]
	if !ok {
		return ErrUnknownCookie
	}
	m.writer.clear(ctx, def)
	return nil
}

func (m *manager) Read(
	ctx *gin.Context,
	t Type,
) (string, error) {

	def, ok := registry[t]
	if !ok {
		return "", ErrUnknownCookie
	}

	return ctx.Cookie(def.Name)
}

// ---- Read helpers ----

func (m *manager) ReadAuth(
	ctx *gin.Context,
) (string, error) {
	return m.Read(ctx, Auth)
}

func (m *manager) ReadSignup(
	ctx *gin.Context,
) (string, error) {
	return m.Read(ctx, Signup)
}

func (m *manager) ReadCSRF(
	ctx *gin.Context,
) (string, error) {
	return m.Read(ctx, CSRF)
}



// ---- Signup helpers ----

// SetSignup allows access to onboarding (OTP verified).
func (m *manager) SetSignup(ctx *gin.Context, value string) error {
	return m.Set(ctx, Signup, value)
}

// ClearSignup removes onboarding access.
func (m *manager) ClearSignup(ctx *gin.Context) {
	_ = m.Clear(ctx, Signup)
}
