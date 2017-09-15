package cacheflight

import (
	"time"
)

// Group is core struct
type Group struct {
	// ...
}

// NewGroup return a new ttl cache group
func NewGroup(cacheExpiration time.Duration) (group *Group) {

	// do something
}

// Do cache
func (g *Group) Do(key string, fn func() (interface{}, error)) (ret interface{}, err error) {

}
