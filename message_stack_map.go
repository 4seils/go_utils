package go_utils

import (
	"sync"
)

type MqStruct struct {
	l       *sync.RWMutex
	m       map[string]*MessageQueue
	Pending chan string // message stack map key
}

const default_q_size = 3

func CreateMessageStackMap(channelCap int) *MqStruct {
	mqs := &MqStruct{
		l:       new(sync.RWMutex),
		m:       make(map[string]*MessageQueue, 0),
		Pending: make(chan string, channelCap),
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
		mq = InitializeMq(default_q_size)
		mqs.m[key] = mq
	}
	mq.Push(0, val)
	mqs.l.Unlock()
	mqs.Pending <- key
}

func (mqs *MqStruct) GetMessageQueue(key string) *MessageQueue {
	mqs.l.RLock()
	mq := mqs.m[key]
	mqs.l.RUnlock()
	return mq
}
