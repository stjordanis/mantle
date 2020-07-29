package queryhandler

import (
	"fmt"
	"log"

	"bytes"
	"reflect"

	"github.com/terra-project/mantle/db"
	"github.com/terra-project/mantle/db/kvindex"
	"github.com/terra-project/mantle/utils"
)

type SeekResolver struct {
	db           db.DB
	kvindexEntry *kvindex.KVIndexEntry
	entityName   string
	indexName    string
	seekKey      []byte
}

// seek resolver
func NewSeekResolver(
	db db.DB,
	kvindexEntry *kvindex.KVIndexEntry,
	entityName,
	indexName string,
	indexOption interface{},
) QueryHandler {
	seekKey, err := kvindexEntry.ResolveKeyType(indexOption)

	// if ResolveKeyType fails, that means
	if err != nil {
		log.Fatalf(
			"Hash index is given but the given option can't be used for index %s, entity=%s, indexOptionType=%s",
			indexName,
			entityName,
			reflect.TypeOf(indexOption).Kind().String(),
		)
		return nil
	}

	return &SeekResolver{
		db:           db,
		kvindexEntry: kvindexEntry,
		entityName:   entityName,
		indexName:    indexName,
		seekKey:      seekKey,
	}
}

func (resolver SeekResolver) Resolve() (interface{}, error) {
	var seekKeyActual = utils.ConcatBytes([]byte(resolver.entityName), []byte(resolver.indexName), []byte(resolver.seekKey))
	it := resolver.db.IndexIterator(
		seekKeyActual,
		seekKeyActual,
		false,
	)

	documentKey := bytes.NewBuffer(nil)
	documentKey.Write([]byte(resolver.entityName))

	if it.Valid() {
		documentKey.Write([]byte(it.DocumentKey()))
		it.Close() // close immediately
	} else {
		return nil, fmt.Errorf(
			"Index does not exist, entityName=%s, indexName=%s, indexKey=%s",
			resolver.entityName,
			resolver.indexName,
			resolver.seekKey,
		)
	}

	return documentKey.Bytes(), nil
}