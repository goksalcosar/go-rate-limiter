package limiter

const (
	LimiterTokenBucket   = "token"
	LimiterSlidingWindow = "sliding"
)

type RateLimiter interface {
	Allow(key string) bool
}