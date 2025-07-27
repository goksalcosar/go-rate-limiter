package limiter

func NewRateLimiter(alg string) RateLimiter {
	switch alg {
	case LimiterSlidingWindow:
		return NewSlidingWindowLimiter(5, 60) // 60 saniyede max 5 istek
	default:
		return NewTokenBucketLimiter(1, 5) // saniyede 1 token, max 5 token
	}
}