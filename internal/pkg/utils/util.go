package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"math/big"

	"bit-project/gateway/config"
)

func CreateRSAKeyFromJWK(n, e string) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(n)
	if err != nil {
		return nil, errors.New("invalid base64 for N")
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(e)
	if err != nil {
		return nil, errors.New("invalid base64 for E")
	}

	nBigInt := new(big.Int).SetBytes(nBytes)
	eBigInt := new(big.Int).SetBytes(eBytes)
	eInt := int(eBigInt.Uint64())

	pubKey := &rsa.PublicKey{
		N: nBigInt,
		E: eInt,
	}
	return pubKey, nil
}

func FindJWKByKID(kid string, jwks []config.Key) (*config.Key, error) {
	for _, key := range jwks {
		if key.Kid == kid {
			return &key, nil
		}
	}
	return nil, errors.New("kid not found")
}
