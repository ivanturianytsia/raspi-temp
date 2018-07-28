package main

import (
	"sync"
	"time"
)

type DataPoint struct {
	Value float32   `json:"value"`
	Time  time.Time `json:"time"`
}

type DataStore struct {
	data DataPoint
	lock *sync.Mutex
}

func NewDataStore() DataStore {
	return DataStore{
		data: DataPoint{},
		lock: &sync.Mutex{},
	}
}

func (store *DataStore) Set(point DataPoint) error {
	store.lock.Lock()
	store.data = point
	store.lock.Unlock()
	return nil
}

func (store *DataStore) Get() DataPoint {
	return store.data
}
