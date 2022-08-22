package storage

import (
	"encoding/json"
	"fmt"

	"github.com/99designs/keyring"
)

type Storage struct {
	Ring keyring.Keyring
}

type Item struct {
	Key   string
	Value string
}

func (r *Storage) CreateEnv(key string) error {

	if r.envExists(key) {
		return fmt.Errorf("env already exists")
	}

	data, _ := json.Marshal([]Item{})

	r.Ring.Set(keyring.Item{
		Key:  key,
		Data: data,
	})

	return nil
}

func (r *Storage) RemoveEnv(key string) error {
	if !r.envExists(key) {
		return fmt.Errorf("env does not exists")
	}

	r.Ring.Remove(key)

	return nil
}

func (r *Storage) GetEnv(key string) ([]Item, error) {
	items, err := r.getItems(key)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Storage) GetEnvs() ([]string, error) {
	return r.Ring.Keys()
}

func (r *Storage) Add(env string, key string, value string) error {

	items, err := r.getItems(env)

	if err != nil {
		return err
	}

	itemIdx := getItemIdxByKey(key, items)

	if itemIdx != -1 {
		items[itemIdx].Value = value
	} else {
		items = append(items, Item{
			Key:   key,
			Value: value,
		})
	}

	data, _ := json.Marshal(items)

	err = r.Ring.Set(keyring.Item{
		Key:  env,
		Data: data,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *Storage) Remove(env string, key string) error {
	items, err := r.getItems(env)

	if err != nil {
		return err
	}

	itemIdx := getItemIdxByKey(key, items)

	if itemIdx != -1 {
		items = append(items[:itemIdx], items[itemIdx+1:]...)
	} else {
		return fmt.Errorf("item not found")
	}

	data, _ := json.Marshal(items)

	err = r.Ring.Set(keyring.Item{
		Key:  env,
		Data: data,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *Storage) envExists(env string) bool {
	_, err := r.Ring.Get(env)
	return err == nil
}

func (r *Storage) getItems(env string) ([]Item, error) {
	v, err := r.Ring.Get(env)

	if err != nil {
		return nil, fmt.Errorf("env not found %w", err)
	}

	var items []Item

	err = json.Unmarshal(v.Data, &items)

	if err != nil {
		return nil, fmt.Errorf("parsing item.Data failed %w", err)
	}

	return items, nil
}

// helpers

func getItemIdxByKey(key string, items []Item) int {
	for idx, item := range items {
		if item.Key == key {
			return idx
		}
	}
	return -1
}
