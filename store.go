package main

import "crypto"

// Data store.

// CertInfo stores certificate information.
type CertInfo struct{}

func (ci *CertInfo) IssuerNameHash(crypto.Hash) []byte {}
func (ci *CertInfo) IssuerKeyHash(crypto.Hash) []byte  {}

// A DataStore maintains the list of certificate revocations.
type DataStore interface {
	Load() ([]CertInfo, error)
}
