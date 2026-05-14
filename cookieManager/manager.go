package cookieManager

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Manager is the public API. Use cookieManager.Auth / Signup / CSRF
// constants as the Type argument — e.g. mgr.Read(ctx, cookieManager.Auth).
type Manager interface {
	Set(ctx *gin.Context, t Type, value string) error
	Clear(ctx *gin.Context, t Type) error
	Read(ctx *gin.Context, t Type) (string, error)
}

var (
	ErrUnknownCookie = errors.New("unknown cookie type")
)

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
