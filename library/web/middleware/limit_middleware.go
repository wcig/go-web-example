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

var (
	rl *RequestLimit
)

func LimitMiddleware(n int64) gin.HandlerFunc {
	rl = &RequestLimit{
		sem: make(chan struct{}, n),
	}
	return func(c *gin.Context) {
		rl.enter()
		defer rl.leave()
		c.Next()
	}
}
