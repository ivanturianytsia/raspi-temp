package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

type Value struct {
	Value    float32   `json:"value"`
	DateTime time.Time `json:"datetime"`
}

type DataPoint struct {
	Data        []Value `json:"data"`
	Unit        string  `json:"unit"`
	Measurement string  `json:"measurement"`
}

type DataStore struct {
	dataPath string
	lock     *sync.Mutex
}

func NewDataStore() (*DataStore, error) {
	path := os.Getenv("DATA_PATH")
	if path == "" {
		return nil, fmt.Errorf("Invalid DATA_PATH env.")
	}
	return &DataStore{
		dataPath: path,
		lock:     &sync.Mutex{},
	}, nil
}

func (store *DataStore) Save(newPoint DataPoint) error {
	filename := path.Join(store.dataPath, fmt.Sprintf("%s.json", time.Now().Local().Format("02Jan2006")))

	store.lock.Lock()
	defer store.lock.Unlock()
	data, err := store.Get(time.Now())
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		file.Close()
	}

	done := false
	for i, v := range data {
		if v.Measurement == newPoint.Measurement && v.Unit == newPoint.Unit {
			done = true
			data[i].Data = append(data[i].Data, newPoint.Data...)
		}
	}
	if !done {
		data = append(data, newPoint)
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, dataBytes, os.ModePerm)
}

func (store *DataStore) Get(datetime time.Time) ([]DataPoint, error) {
	filename := path.Join(store.dataPath, fmt.Sprintf("%s.json", datetime.Local().Format("02Jan2006")))
	data := []DataPoint{}
	if _, err := os.Stat(filename); err != nil {
		return data, err
	}
	dataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return data, err
	}
	return data, nil
}
