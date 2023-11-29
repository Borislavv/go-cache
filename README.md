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
const UserExampleID = 1

type User struct {
    ID int
}

func main() {
	ctx, _ := context.WithCancel(context.Background())

	storage   := NewMapCacheStorage(ctx)
	displacer := NewCacheDisplacer(ctx, time.Second*1)
	cache     := NewCache(storage, displacer)

	cacheKey  := fmt.Sprintf("userID_%v", UserExampleID)

	userInterface, _ := cache.Get(cacheKey, func(item CacheItem) (data interface{}, err error) {
		item.SetTTL(time.Hour)
		return fetchUserByID(UserExampleID), nil
	})

	cachedUser, _ := userInterface.(*User)

	fmt.Printf("userID is %d", cachedUser.ID)
	
	// output:
	// 	user fetched from storage
	//	userID is 1

	userInterface, _ = cache.Get(cacheKey, func(item CacheItem) (data interface{}, err error) {
		item.SetTTL(time.Hour)
		return fetchUserByID(UserExampleID), nil
	})

	cachedUser, _ = userInterface.(*User)

	fmt.Println(cachedUser.ID)
	
	// output:
	//	userID is 1
}

func fetchUserByID(id int) *User {
	fmt.Print("user fetched from storage")
	return &User{ID: id}
}
```
