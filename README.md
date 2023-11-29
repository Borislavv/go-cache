# go-cache

## Cache implementation for golang

Interface:
```
type Cacher interface {
	Get(key string, fn func(CacheItem) (data interface{}, err error)) (data interface{}, err error)
	Delete(key string)
}

type CacheItem interface {
	SetTTL(ttl time.Duration)
}
```

Create a new instance of cache:
```
ctx, _    := context.WithCancel(context.Background())
storage   := cache.NewMapCacheStorage(ctx)
displacer := cache.NewCacheDisplacer(ctx, time.Second*1)
cache     := cache.NewCache(storage, displacer)
```

Usage: 
```
package main

import (
	"context"
	"fmt"
	"github.com/Borislavv/go-cache"
	"time"
)

const UserExampleID = 1

type User struct {
	ID int
}

func main() {
	ctx, _ := context.WithCancel(context.Background())

	storage 	:= cache.NewMapCacheStorage(ctx)
	displacer 	:= cache.NewCacheDisplacer(ctx, time.Second*1)
	cacher		:= cache.NewCache(storage, displacer)

	cacheKey 	:= fmt.Sprintf("userID_%v", UserExampleID)

	userInterface, _ := cacher.Get(cacheKey, func(item cache.CacheItem) (data interface{}, err error) {
		item.SetTTL(time.Hour)
		return fetchUserByID(UserExampleID), nil
	})

	cachedUser, _ := userInterface.(*User)

	fmt.Printf("userID is %d\n", cachedUser.ID)

	// output:
	// 	user fetched from storage
	//	userID is 1

	userInterface, _ = cacher.Get(cacheKey, func(item cache.CacheItem) (data interface{}, err error) {
		item.SetTTL(time.Hour)
		return fetchUserByID(UserExampleID), nil
	})

	cachedUser, _ = userInterface.(*User)

	fmt.Printf("userID is %d\n", cachedUser.ID)

	// output:
	//	userID is 1
}

func fetchUserByID(id int) *User {
	fmt.Println("user fetched from storage")
	return &User{ID: id}
}

// Total output:
// 	user fetched from storage
//	userID is 1
//	userID is 1
```
