package keystore

import (
	"errors"

	"github.com/filecoin-project/lotus/chain/types"
)

// MemKeystore is an in memory keystore implementation that is not persisted to
// any backing storage.
type MemKeystore struct {
	keys map[string]types.KeyInfo
}

// NewMemKeystore creates a MemKeystore.
func NewMemKeystore() *MemKeystore {
	return &MemKeystore{make(map[string]types.KeyInfo)}
}

// Has return whether or not a key exists in the Keystore
func (mk *MemKeystore) Has(name string) (bool, error) {
	_, ok := mk.keys[name]
	return ok, nil
}

// Put store a key in the Keystore
func (mk *MemKeystore) Put(name string, k types.KeyInfo) error {
	if name == "" {
		return errors.New("key name must be at least one character")
	}

	_, ok := mk.keys[name]
	if ok {
		return ErrKeyExists
	}

	mk.keys[name] = k
	return nil
}

// Get retrieve a key from the Keystore
func (mk *MemKeystore) Get(name string) (types.KeyInfo, error) {
	k, ok := mk.keys[name]
	if !ok {
		return types.KeyInfo{}, ErrNoSuchKey
	}

	return k, nil
}

// Delete remove a key from the Keystore
func (mk *MemKeystore) Delete(name string) error {
	delete(mk.keys, name)
	return nil
}

// List return a list of key identifier
func (mk *MemKeystore) List() ([]string, error) {
	out := make([]string, 0, len(mk.keys))
	for k := range mk.keys {
		out = append(out, k)
	}
	return out, nil
}
