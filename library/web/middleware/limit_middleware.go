package middleware

import "github.com/gin-gonic/gin"

type RequestLimit struct {
	sem chan struct{}
}

func (rl *RequestLimit) enter() {
	rl.sem <- struct{}{}
}

func (rl *RequestLimit) leave() {
	<-rl.sem
}

func NewRequestLimit(n int64) *RequestLimit {
	return &RequestLimit{
		sem: make(chan struct{}, n),
	}
}

var (
	rl *RequestLimit
)

func LimitMiddleware(n int64) gin.HandlerFunc {
	rl = NewRequestLimit(n)
	return func(c *gin.Context) {
		rl.enter()
		defer rl.leave()
		c.Next()
	}
}
