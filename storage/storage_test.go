package storage_test

import (
	"testing"

	"github.com/99designs/keyring"
	"github.com/janstuemmel/env-vault/storage"
)

func TestCreateEnv(t *testing.T) {
	ring := keyring.NewArrayKeyring([]keyring.Item{})
	storage := storage.Storage{Ring: ring}

	storage.CreateEnv("dummy")

	v, _ := ring.Get("dummy")
	assert(t, "dummy", v.Key)
	assert(t, "[]", string(v.Data))

	_, err := ring.Get("i-do-not-exist")
	assert(t, "The specified item could not be found in the keyring", err.Error())
}

func TestAddSecretVarAndUpdate(t *testing.T) {
	ring := keyring.NewArrayKeyring([]keyring.Item{
		{Key: "dummy", Data: []byte("[]")},
	})
	storage := storage.Storage{Ring: ring}

	storage.Add("dummy", "foo", "bar")

	v, _ := ring.Get("dummy")
	assert(t, "dummy", v.Key)
	assert(t, `[{"Key":"foo","Value":"bar"}]`, string(v.Data))

	storage.Add("dummy", "foo", "baz")

	v, _ = ring.Get("dummy")
	assert(t, "dummy", v.Key)
	assert(t, `[{"Key":"foo","Value":"baz"}]`, string(v.Data))
}

func assert(t *testing.T, want interface{}, have interface{}) {

	// mark as test helper
	t.Helper()

	// throw error
	if want != have {
		t.Errorf("Assertion failed for %s\n\twant:\t%+v\n\thave:\t%+v", t.Name(), want, have)
	}
}
