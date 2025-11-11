package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	idCounter int
	idMutex   sync.Mutex
)

// GenerateID creates a unique ID for entities
func GenerateID() string {
	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter++
	return fmt.Sprintf("%d", idCounter)
}

// GenerateIDWithTimestamp creates an ID with timestamp for uniqueness
func GenerateIDWithTimestamp() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
}

// IsValidEntityName checks if entity name is valid
func IsValidEntityName(name string) bool {
	return len(name) > 0 && len(name) <= 255
}

// IsValidDescription checks if description is valid
func IsValidDescription(desc string) bool {
	return len(desc) <= 1000
}
