package client

import "sync"

type MarketDataControl struct {
	table []MarketDataAction
	free  *MarketDataAction
	mu    sync.Mutex
}

func NewMarketDataControl(size int) {
	if size < 100 {
		size = 100
	}
	c := &MarketDataControl{
		table: make([]MarketDataAction, size),
	}
	for i := 0; i < len(c.table); i++ {
		a := &c.table[i]
		a.control = c
		a.id = int32(i)
	}
	for i := 0; i < len(c.table); i++ {
		a := &c.table[i]
		if i > 0 {
			a.prev = &c.table[i-1]
		}
		if i < len(c.table)-1 {
			a.next = &c.table[i+1]
		}
	}
	c.free = &c.table[0]
}

func (m *MarketDataControl) Action(id uint32) *MarketDataAction {
	if id < uint32(len(m.table)) {
		return &m.table[id]
	}
	return nil
}

func (m *MarketDataControl) New() *MarketDataAction {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.free == nil {
		return nil
	}
	a := m.free
	if a.next != nil {
		m.free = a.next
		m.free.prev = nil
	}
	return a
}

func (m *MarketDataControl) push(a *MarketDataAction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.free == nil {
		m.free = a
		return
	}
	m.free.prev = a
	a.next = m.free
	m.free = a
}

type MarketDataAction struct {
	control *MarketDataControl
	id      int32
	prev    *MarketDataAction
	next    *MarketDataAction
	mu      sync.Mutex
}

func (m *MarketDataAction) Id() int32 {
	return m.id
}

func (m *MarketDataAction) Close() error {
	m.control.push(m)
	return nil
}
