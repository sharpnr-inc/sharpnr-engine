package cookieManager

import "github.com/gin-gonic/gin"

type Manager interface {
	Set(ctx *gin.Context, t Type, value string) error
	Clear(ctx *gin.Context, t Type) error
	Read(ctx *gin.Context, t Type) (string, error)
	ReadAuth(ctx *gin.Context) (string, error)
	ReadSignup(ctx *gin.Context) (string, error)
	ReadCSRF(ctx *gin.Context) (string, error)
	SetSignup(ctx *gin.Context, value string) error
	ClearSignup(ctx *gin.Context)
}

type manager struct {
	defs map[Type]Definition
}

func New(defs map[Type]Definition) Manager {
	return &manager{defs: defs}
}

func (m *manager) resolve(t Type) (Definition, error) {
	def, ok := m.defs[t]
	if !ok {
		return Definition{}, ErrUnknownCookie
	}
	return def, nil
}

func (m *manager) Set(ctx *gin.Context, t Type, value string) error {
	def, err := m.resolve(t)
	if err != nil {
		return err
	}
	write(ctx, def, value)
	return nil
}

func (m *manager) Clear(ctx *gin.Context, t Type) error {
	def, err := m.resolve(t)
	if err != nil {
		return err
	}
	clear(ctx, def)
	return nil
}

func (m *manager) Read(ctx *gin.Context, t Type) (string, error) {
	def, err := m.resolve(t)
	if err != nil {
		return "", err
	}
	return ctx.Cookie(def.Name)
}

func (m *manager) ReadAuth(ctx *gin.Context) (string, error)   { return m.Read(ctx, Auth) }
func (m *manager) ReadSignup(ctx *gin.Context) (string, error) { return m.Read(ctx, Signup) }
func (m *manager) ReadCSRF(ctx *gin.Context) (string, error)   { return m.Read(ctx, CSRF) }

func (m *manager) SetSignup(ctx *gin.Context, value string) error { return m.Set(ctx, Signup, value) }
func (m *manager) ClearSignup(ctx *gin.Context)                   { _ = m.Clear(ctx, Signup) }