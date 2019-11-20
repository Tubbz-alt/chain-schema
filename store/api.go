package store

import (
	abci "github.com/tendermint/tendermint/types"
	"io"
	"context"
	"github.com/iov-one/weave"
)

type ReadonlyContext interface {
	context.Context
	ReadonlyKVStore(prefix []byte) weave.ReadOnlyKVStore
	BlockHeader() abci.Header
}

type MutableContext interface {
	ReadonlyContext
	KVStore(prefix []byte) weave.KVStore
}

type Index interface {
	Has(ctx context.Context, key []byte) (bool, error)
	Get(ctx context.Context, key []byte) (Iterator, error)
	PrefixScan(ctx context.Context, start []byte, end []byte) (Iterator, error)
	ReversePrefixScan(ctx context.Context, start []byte, end []byte) (Iterator, error)
}

type UniqueIndex interface {
	Index
	GetOne(ctx context.Context, indexKey []byte, dest interface{}) (primaryKey []byte, error error)
}

type UInt64Index interface {
	Has(ctx context.Context, key uint64) (bool, error)
	Get(ctx context.Context, key uint64) (Iterator, error)
	PrefixScan(ctx context.Context, start uint64, end uint64) (Iterator, error)
	ReversePrefixScan(ctx context.Context, start uint64, end uint64) (Iterator, error)
}

type Indexer interface {
	//DoIndex(store sdk.KVStore, rowId uint64, key []byte, value interface{}) error
	//BuildIndex(storeKey sdk.StoreKey, prefix []byte, modelGetter func(ctx Context, rowId uint64, dest interface{}) (key []byte, err error)) Index
}

type BucketBuilder interface {
	CreateIndex(prefix []byte, indexer Indexer) Index
	Build() BucketBase
}

// BucketBase provides methods shared by all buckets
type BucketBase interface {
	UniqueIndex
	// Delete deletes the value at the given key
	Delete(ctx context.Context, key []byte) error
}

// ExternalKeyTable defines a bucket where the key is stored externally to the value object
type ExternalKeyTable interface {
	BucketBase
	// Save saves the given key value pair
	Save(ctx context.Context, key []byte, value interface{}) error
}

type HasID interface {
	ID() []byte
}

// NaturalKeyTable defines a bucket where all values implement HasID and the key is stored it the value and
// returned by the HasID method
type NaturalKeyTable interface {
	BucketBase
	// Save saves the value passed in
	Save(ctx context.Context, value HasID) error
}

type AutoUInt64Table interface {
	Has(ctx context.Context, key uint64) (bool, error)
	Get(ctx context.Context, key uint64) (Iterator, error)
	PrefixScan(ctx context.Context, start uint64, end uint64) (Iterator, error)
	ReversePrefixScan(ctx context.Context, start uint64, end uint64) (Iterator, error)
	Save(ctx context.Context, key []byte, value interface{}) error
}

// AutoKeyTable specifies a bucket where keys are generated via an auto-incremented interger
type AutoKeyTable interface {
	ExternalKeyTable

	// Create auto-generates key
	Create(ctx context.Context, value interface{}) ([]byte, error)
}

//Iterator allows iteration through a sequence of key value pairs
type Iterator interface {
	// LoadNext loads the next value in the sequence into the pointer passed as dest and returns the key. If there
	// are no more items an error is returned
	LoadNext(dest interface{}) (key []byte, err error)
	// Close releases the iterator and should be called at the end of iteration
	io.Closer
}

type UInt64Iterator interface {
	// LoadNext loads the next value in the sequence into the pointer passed as dest and returns the key. If there
	// are no more items an error is returned
	LoadNext(dest interface{}) (key uint64, err error)
	// Close releases the iterator and should be called at the end of iteration
	io.Closer
}

type TableBuilder interface {
	RegisterIndexer(prefix byte, indexer Indexer)
}

type NaturalKeyTableBuilder interface {
	TableBuilder
	Build() NaturalKeyTable
}

type AutoUInt64TableBuilder interface {
	RegisterIndexer(prefix byte, indexer Indexer)
	Build() AutoUInt64Table
}
