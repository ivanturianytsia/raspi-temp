package main

import (
	"sync"
	"time"
)

type Value struct {
	Value    float32   `json:"value"`
	DateTime time.Time `json:"datetime"`
}

type DataPoint struct {
	Data        Value  `json:"data"`
	Unit        string `json:"unit"`
	Measurement string `json:"measurement"`
}

type DataStore struct {
	data []DataPoint
	lock *sync.Mutex
}

func NewDataStore() DataStore {
	return DataStore{
		data: []DataPoint{},
		lock: &sync.Mutex{},
	}
}

func (store *DataStore) Set(point DataPoint) error {
	store.lock.Lock()
	store.data = append(store.data, point)
	store.lock.Unlock()
	return nil
}

func (store *DataStore) Get() []DataPoint {
	return store.data
}
