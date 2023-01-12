package gostorage

import (
	"github.com/link-yundi/ytools/ylog"
	"sync"
)

/**
@Since 2022-12-02 13:34
@Author: Huang
@Description: save for multiple goroutines
**/

type MapL1 struct {
	Lock sync.RWMutex
	Map  map[interface{}]interface{}
}

func NewMapL1() MapL1 {
	return MapL1{
		Lock: sync.RWMutex{},
		Map:  make(map[interface{}]interface{}, 0),
	}
}

type MapL2 struct {
	Lock sync.RWMutex
	Map  map[interface{}]map[interface{}]interface{}
}

func NewMapL2() MapL2 {
	return MapL2{
		Lock: sync.RWMutex{},
		Map:  make(map[interface{}]map[interface{}]interface{}, 0),
	}
}

func (m *MapL1) Set(name interface{}, value interface{}) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	if len(m.Map) == 0 {
		m.Map = make(map[interface{}]interface{}, 0)
	}

	m.Map[name] = value
}

func (m *MapL1) Get(name interface{}) (interface{}, bool) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	if val, ok := m.Map[name]; ok {
		return val, true
	}

	return nil, false
}

func (m *MapL1) Delete(name interface{}) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	delete(m.Map, name)
}

func (m *MapL1) Clear() {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	m.Map = make(map[interface{}]interface{}, 0)
}

func (m *MapL1) Size() int {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	return len(m.Map)
}

func (m *MapL1) Copy() map[interface{}]interface{} {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	_copy := make(map[interface{}]interface{}, 0)

	for k, v := range m.Map {
		_copy[k] = v
	}

	return _copy
}

func (m2 *MapL2) Set(key1 interface{}, key2 interface{}, val interface{}) {
	m2.Lock.Lock()
	defer m2.Lock.Unlock()

	if len(m2.Map) == 0 {
		m2.Map = make(map[interface{}]map[interface{}]interface{}, 0)
	}

	if len(m2.Map[key1]) == 0 {
		m2.Map[key1] = make(map[interface{}]interface{}, 0)
	}

	m2.Map[key1][key2] = val
}

func (m2 *MapL2) GetL2(key1 interface{}, key2 interface{}) (interface{}, bool) {
	m2.Lock.RLock()
	defer m2.Lock.RUnlock()

	if val, ok := m2.Map[key1][key2]; ok {
		return val, true
	}

	return nil, false
}

func (m2 *MapL2) GetL1(key1 interface{}) (map[interface{}]interface{}, bool) {
	m2.Lock.RLock()
	defer m2.Lock.RUnlock()

	mLists, ok := m2.Map[key1]
	if !ok {
		return nil, false
	}

	_copy := make(map[interface{}]interface{}, 0)

	for k, v := range mLists {
		_copy[k] = v
	}

	return _copy, true
}

func (m2 *MapL2) Delete(keys ...interface{}) {
	m2.Lock.Lock()
	defer m2.Lock.Unlock()

	switch len(keys) {
	case 1:
		delete(m2.Map, keys[0])
	case 2:
		delete(m2.Map[keys[0]], keys[1])
	default:
		ylog.Warnf("Params %v out of Bound, the mapL2 have no change", keys)
	}
}

func (m2 *MapL2) Clear() {
	m2.Lock.Lock()
	defer m2.Lock.Unlock()

	m2.Map = make(map[interface{}]map[interface{}]interface{}, 0)
}

func (m2 *MapL2) Size(key ...interface{}) int {
	m2.Lock.RLock()
	defer m2.Lock.RUnlock()

	if key == nil {
		return len(m2.Map)
	} else {
		return len(m2.Map[key[0]])
	}
}

func (m2 *MapL2) Copy() map[interface{}]map[interface{}]interface{} {
	m2.Lock.RLock()
	defer m2.Lock.RUnlock()

	_copy := make(map[interface{}]map[interface{}]interface{}, 0)
	for k, v := range m2.Map {
		_copyL2 := make(map[interface{}]interface{}, 0)
		for key, val := range v {
			_copyL2[key] = val
		}
		_copy[k] = _copyL2
	}

	return _copy
}
