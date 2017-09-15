package cacheflight

import (
	"sync"
	"time"
)

// Group is core struct
type Group struct {
	// ...
	Expiration time.Duration
	Cache      map[string]*CacheVal
	Mu         sync.Mutex
}

type CacheVal struct {
	CreatTime int64
	Ret       interface{}
	Err       error
}

// NewGroup return a new ttl cache group
func NewGroup(cacheExpiration time.Duration) (group *Group) {
	g := &Group{}
	g.Cache = make(map[string]*CacheVal)
	g.Expiration = cacheExpiration
	return g
	// do something
}

// Do cache
func (g *Group) Do(key string, fn func() (interface{}, error)) (ret interface{}, err error) {
	now := time.Now().UnixNano()
	//fmt.Printf("[Do] now: %d\n", now)
	g.Mu.Lock()
	defer g.Mu.Unlock()
	if g.Cache[key] == nil {
		//fmt.Printf("[Do]new a cache of [%s]\n", key)
		g.Cache[key] = new(CacheVal)
	} else {
	}
	//fmt.Printf("[Do] time pass: [%s]%d\n", key, now-g.Cache[key].CreatTime)
	//fmt.Printf("[Do] expiration: %d\n", int64(g.Expiration))
	if now-g.Cache[key].CreatTime > int64(g.Expiration) {
		//fmt.Printf("[Do] update cache of [%s] \n", key)
		//g.Mu.Lock()
		g.Cache[key].CreatTime = now
		g.Cache[key].Ret, g.Cache[key].Err = fn()
		//g.Mu.Unlock()
	} else {
		//fmt.Printf("[Do] use the previous result\n")

	}

	return g.Cache[key].Ret, g.Cache[key].Err
}
