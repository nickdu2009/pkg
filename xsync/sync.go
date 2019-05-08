package xsync

import "sync"

func WithLock(l sync.Locker, f func()) {
	l.Lock()
	defer l.Unlock()
	f()
}
