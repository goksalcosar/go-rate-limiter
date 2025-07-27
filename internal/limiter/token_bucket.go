package limiter

import (
	"sync"
	"time"
)

// 👇 Her IP için ayrı bir bucket olacak
type bucket struct {
    tokens         float64   // Şu anki token sayısı
    lastUpdateTime time.Time // En son token üretimi zamanı
}

// 👇 Tüm IP'lerin bucket'larını tutan limiter
type TokenBucket struct {
    rate      float64               // saniyede kaç token üretilecek
    capacity  float64               // maksimum token sayısı
    buckets   map[string]*bucket    // IP → bucket ilişkisi
    mutex     sync.Mutex            // eşzamanlı erişim için kilit
}

func NewTokenBucketLimiter(rate float64, capacity float64) *TokenBucket {
    return &TokenBucket{
        rate:     rate,
        capacity: capacity,
        buckets:  make(map[string]*bucket),
    }
}

func (tokeBucket *TokenBucket) Allow(IP string) bool {
	tokeBucket.mutex.Lock()
    defer tokeBucket.mutex.Unlock()

	now := time.Now()

	buckets, exists := tokeBucket.buckets[IP]

	if !exists {
		buckets = &bucket{
			tokens: tokeBucket.capacity,
			lastUpdateTime: now,
		}

		tokeBucket.buckets[IP] = buckets
	}

	elapsed := now.Sub(buckets.lastUpdateTime).Seconds()

	newTokens := elapsed * tokeBucket.rate

	buckets.tokens = min(tokeBucket.capacity, buckets.tokens + newTokens)
	buckets.lastUpdateTime = now

	if buckets.tokens >= 1 {
		buckets.tokens--

		return true
	}

	return false
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
	
    return b
}