package db

import (
	tmdb "github.com/tendermint/tm-db"
)

type DB interface {
	Get(key []byte) ([]byte, error)
	Set(key, data []byte) error
	Delete(key []byte) error
	Iterator(start []byte, reverse bool) Iterator
	IndexIterator(start []byte, reverse bool) Iterator
	Batch() Batch
	GetCosmosAdapter() tmdb.DB
}

type Iterator interface {
	Close()
	Valid(prefix []byte) bool
	Next()
	Key() []byte
	DocumentKey() []byte
	Value() []byte
}

type Batch interface {
	Set(key, data []byte) error
	Delete(key []byte) error
	Flush() error
	Close()
}