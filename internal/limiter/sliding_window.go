package limiter

import (
	"sync"
	"time"
)

type SlidingWindow struct {
    limit     int                          // Maksimum istek sayısı
    window    time.Duration                // Zaman penceresi (örn: 60 saniye)
    requests  map[string][]time.Time       // IP -> yapılan istek zamanları
    mutex     sync.Mutex                   // Veri güvenliği için
}

func NewSlidingWindowLimiter(limit int, windowSeconds int) *SlidingWindow {
    return &SlidingWindow{
        limit:    limit,
        window:   time.Duration(windowSeconds) * time.Second,
        requests: make(map[string][]time.Time),
    }
}

func (sw *SlidingWindow) Allow(IP string) bool {
    sw.mutex.Lock()
    defer sw.mutex.Unlock()

    now := time.Now()

	//Zamanda geriye gitmek için "-" kullanılır.
    windowStart := now.Add(-sw.window)

    // IP'ye ait tüm istek zamanlarını al
    times := sw.requests[IP]

    // Pencere dışındaki (eski) istekleri temizle
    var recent []time.Time

    for _, time := range times {
        if time.After(windowStart) {
            recent = append(recent, time)
        }
    }

    // Eğer limit aşılmışsa reddet
    if len(recent) >= sw.limit {
        return false
    }

    // Yeni isteği kaydet
    recent = append(recent, now)
    sw.requests[IP] = recent

    return true
}
