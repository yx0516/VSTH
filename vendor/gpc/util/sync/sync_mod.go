package sync

import (
	"sync"
)

//-----------------------------------------------------------------------------------------------------------//

// 实现，每 n 次就返回一次 True 的计数器
type SyncMod struct {
	lock  sync.Mutex
	count int
	mod   int
}

func NewSyncMod(mod int) *SyncMod {
	return &SyncMod{
		count: 0,
		mod:   mod,
	}
}

//-----------------------------------------------------------------------------------------------------------//

func (self *SyncMod) SetMod(mod int) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.mod = mod
}

func (self *SyncMod) Do() bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	if self.mod <= 0 {
		return false
	} else {
		self.count++
		if (self.count % self.mod) == 0 {
			self.count = 0
			return true
		} else {
			return false
		}
	}
}

//-----------------------------------------------------------------------------------------------------------//
