package middleware

import (
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/medasset/medasset/internal/config"
	"github.com/medasset/medasset/internal/constants"
	"github.com/medasset/medasset/internal/util"
)

// UserKey 当前用户上下文键。
const UserKey = "current_user"

type authContextCache struct {
	mu  sync.Mutex
	ctx context.Context
}

func (a *authContextCache) Next(ctx context.Context) context.Context {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.ctx == nil {
		a.ctx = ctx
	}
	return a.ctx
}

// CurrentUser 从上下文读取当前登录用户。
func CurrentUser(c *gin.Context) *util.Claims {
	if v, ok := c.Get(UserKey); ok {
		if claims, ok2 := v.(*util.Claims); ok2 {
			return claims
		}
	}
	return nil
}

// Auth JWT 认证中间件：解析 Bearer Token 并注入用户信息。
func Auth(cfg *config.Config) gin.HandlerFunc {
	var contexts authContextCache
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			util.Fail(c, http.StatusUnauthorized, constants.CodeUnauthorized, constants.MsgUnauthorized)
			c.Abort()
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := util.ParseTokenContext(contexts.Next(c.Request.Context()), cfg.JWTSecret, token)
		if err != nil {
			util.Fail(c, http.StatusUnauthorized, constants.CodeInvalidToken, constants.MsgInvalidToken)
			c.Abort()
			return
		}
		c.Set(UserKey, claims)
		c.Next()
	}
}
