package distributed_lock

import "github.com/go-redis/redis/v8"

// todo: complete the implementation of MyRedLock.

// MyRedLock this is my own implementation of RedLock.
// Check if it's correct by comparing with https://github.com/go-redsync/redsync.
type MyRedLock struct {
	Pool []*redis.Client
}

func (m *MyRedLock) Lock() {

}

func (m *MyRedLock) UnLock() {

}
