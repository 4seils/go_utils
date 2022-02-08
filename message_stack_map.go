package go_utils

import (
	"sync"
)

type MqStruct struct {
	l       *sync.RWMutex
	m       map[string]*MessageQueue
	Pending chan string // message stack map key
	cap     int         // channel and MessageQueue size
}

func CreateMessageStackMap(cap int) *MqStruct {
	mqs := &MqStruct{
		l:       new(sync.RWMutex),
		m:       make(map[string]*MessageQueue, 0),
		Pending: make(chan string, cap),
		cap:     cap,
	}
	return mqs
}

func (mqs *MqStruct) AddToMessageStackMap(key string, cap int) {
	mqs.l.Lock()
	mqs.m[key] = InitializeMq(cap)
	mqs.l.Unlock()
}

func (mqs *MqStruct) SetMessageQueue(key string, val interface{}) {
	mqs.l.Lock()
	mq, ok := mqs.m[key]
	if !ok {
		mq = InitializeMq(mqs.cap)
		mqs.m[key] = mq
	}
	mq.Push(0, val)
	mqs.l.Unlock()
	mqs.Pending <- key
}

func (mqs *MqStruct) GetMessageQueue(key string) (*MessageQueue, bool) {
	mqs.l.RLock()
	mq, valid := mqs.m[key]
	mqs.l.RUnlock()
	return mq, valid
}
