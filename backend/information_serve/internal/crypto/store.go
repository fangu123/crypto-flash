package crypto

import (
	"crypto/rsa"
	"sync"
)

type PrivateKeyStore struct {
	mu  sync.RWMutex
	key *rsa.PrivateKey
}
type PublicKeyStore struct {
	key *rsa.PublicKey
}

var PrivateStore = &PrivateKeyStore{}

func (s *PrivateKeyStore) Set(key *rsa.PrivateKey) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.key = key
}

func (s *PrivateKeyStore) Get() *rsa.PrivateKey {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.key
}

var PublicStore = &PublicKeyStore{}

func (s *PublicKeyStore) Set(key *rsa.PublicKey) {
	s.key = key
}

func (s *PublicKeyStore) Get() *rsa.PublicKey {
	return s.key
}
