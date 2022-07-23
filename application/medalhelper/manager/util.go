package manager

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomString(length int) (sink string) {
	source := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	rand.Shuffle(len(source), func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})
	for i := 0; i < length; i++ {
		sink += source[i]
	}
	return
}

// GetTimestamp can obtain current ts
func GetTimestamp() string {
	return fmt.Sprintf("%d", time.Now().UnixMilli())
}
