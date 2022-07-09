package middleware

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/Hank-Kuo/personal-web-backend/pkg/response"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type RateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

func NewRateLimit(r rate.Limit, b int) *RateLimiter {
	i := &RateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
	return i
}

func (i *RateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

func (i *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()

	return limiter
}

func (m *Middleware) RateLimit() func(c *gin.Context) {
	rateLimiter := rate.Every(time.Second)
	limiter := NewRateLimit(rateLimiter, m.cfg.Server.RateLimitPerSec)
	return func(c *gin.Context) {
		host := c.Request.Host
		limiter := limiter.GetLimiter(host)
		if !limiter.Allow() {
			response.Fail(utils.HttpError{Status: http.StatusTooManyRequests, Message: "too many request", Detail: errors.New("too many request")}, m.logger).ToJSON(c)
			return
		}

		c.Next()
	}
}
