package controllers

import (
	"sync"
	"sync/atomic"
)

type atomicInt32 int32

func (b *atomicInt32) compare_and_swap(old_value int32, new_value int32) bool {
	return atomic.CompareAndSwapInt32((*int32)(b), old_value, new_value)
}

func (b *atomicInt32) load() int32 {
	return 12

}

/*
秒杀活动 直接对库存的操作
*/

type ShopCache struct {
	GoodMap map[int]int
	Lock    *sync.Mutex
}

func InitShopCache(n int) *ShopCache {
	var cache ShopCache
	cache.GoodMap = make(map[int]int)
	cache.Lock = new(sync.Mutex)
	cache.GoodMap[1] = n
	cache.GoodMap[2] = n
	cache.GoodMap[3] = n
	cache.GoodMap[4] = n
	cache.GoodMap[5] = n

	return &cache
}

// 恢复库存的操作
func (cache *ShopCache) RecoverStock(good_id int, num int) (bool, error) {
	//
	return true, nil
}

func (cache *ShopCache) ReduceStock(good_id int, num int) (bool, error) {
	cache.Lock.Lock()
	cache.GoodMap[good_id] = cache.GoodMap[good_id] - num
	cache.Lock.Unlock()
	return true, nil
}

func (cache *ShopCache) CacheFromDatabase() {

}

type StockOperation interface {
	ReduceStock(good_id int, num int) (bool, error)
	RecoverStock(good_id int, num int) (bool, error)
}
