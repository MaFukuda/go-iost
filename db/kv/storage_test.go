package kv

import (
	"crypto/rand"
	"os/exec"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	DBPATH = "db"
)

type StorageTestSuite struct {
	suite.Suite
	storage *Storage
	t       StorageType
}

func (suite *StorageTestSuite) SetupTest() {
	storage, err := NewStorage(DBPATH, suite.t)
	suite.Require().Nil(err)
	suite.storage = storage
	suite.storage.Put([]byte("key01"), []byte("value01"))
	suite.storage.Put([]byte("key02"), []byte("value02"))
	suite.storage.Put([]byte("key03"), []byte("value03"))
	suite.storage.Put([]byte("key04"), []byte("value04"))
	suite.storage.Put([]byte("key05"), []byte("value05"))
	suite.storage.Put([]byte("iost01"), []byte("value06"))
	suite.storage.Put([]byte("iost02"), []byte("value07"))
	suite.storage.Put([]byte("iost03"), []byte("value08"))
	suite.storage.Put([]byte("iost04"), []byte("value09"))
	suite.storage.Put([]byte("iost05"), []byte("value10"))
}

func (suite *StorageTestSuite) TestGet() {
	var value []byte
	var err error
	value, err = suite.storage.Get([]byte("key00"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)
	value, err = suite.storage.Get([]byte("key01"))
	suite.Nil(err)
	suite.Equal([]byte("value01"), value)
	value, err = suite.storage.Get([]byte("key05"))
	suite.Nil(err)
	suite.Equal([]byte("value05"), value)
}

func (suite *StorageTestSuite) TestPut() {
	var value []byte
	var err error
	value, err = suite.storage.Get([]byte("key06"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)
	value, err = suite.storage.Get([]byte("key07"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)

	err = suite.storage.Put([]byte("key07"), []byte("value07"))
	suite.Nil(err)

	value, err = suite.storage.Get([]byte("key06"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)
	value, err = suite.storage.Get([]byte("key07"))
	suite.Nil(err)
	suite.Equal([]byte("value07"), value)
}

func (suite *StorageTestSuite) TestDelete() {
	var value []byte
	var err error
	value, err = suite.storage.Get([]byte("key04"))
	suite.Nil(err)
	suite.Equal([]byte("value04"), value)
	value, err = suite.storage.Get([]byte("key05"))
	suite.Nil(err)
	suite.Equal([]byte("value05"), value)

	err = suite.storage.Delete([]byte("key04"))
	suite.Nil(err)

	value, err = suite.storage.Get([]byte("key04"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)
	value, err = suite.storage.Get([]byte("key05"))
	suite.Nil(err)
	suite.Equal([]byte("value05"), value)
}

func (suite *StorageTestSuite) TestKeys() {
	var keys [][]byte
	var err error

	keys, err = suite.storage.Keys([]byte("key"))
	suite.Nil(err)
	suite.ElementsMatch(
		[][]byte{
			[]byte("key01"),
			[]byte("key02"),
			[]byte("key03"),
			[]byte("key04"),
			[]byte("key05"),
		},
		keys,
	)
	keys, err = suite.storage.Keys([]byte("iost"))
	suite.Nil(err)
	suite.ElementsMatch(
		[][]byte{
			[]byte("iost01"),
			[]byte("iost02"),
			[]byte("iost03"),
			[]byte("iost04"),
			[]byte("iost05"),
		},
		keys,
	)
}

func (suite *StorageTestSuite) TestBatch() {
	var value []byte
	var err error

	value, err = suite.storage.Get([]byte("key04"))
	suite.Nil(err)
	suite.Equal([]byte("value04"), value)
	value, err = suite.storage.Get([]byte("key05"))
	suite.Nil(err)
	suite.Equal([]byte("value05"), value)
	value, err = suite.storage.Get([]byte("key06"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)

	err = suite.storage.BeginBatch()
	suite.Nil(err)
	err = suite.storage.BeginBatch()
	suite.NotNil(err)

	err = suite.storage.Delete([]byte("key04"))
	suite.Nil(err)
	err = suite.storage.Put([]byte("key06"), []byte("value06"))
	suite.Nil(err)

	err = suite.storage.CommitBatch()
	suite.Nil(err)
	err = suite.storage.CommitBatch()
	suite.NotNil(err)

	value, err = suite.storage.Get([]byte("key04"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)
	value, err = suite.storage.Get([]byte("key05"))
	suite.Nil(err)
	suite.Equal([]byte("value05"), value)
	value, err = suite.storage.Get([]byte("key06"))
	suite.Nil(err)
	suite.Equal([]byte("value06"), value)
}

func (suite *StorageTestSuite) TestRecover() {
	var value []byte
	var err error

	value, err = suite.storage.Get([]byte("key04"))
	suite.Nil(err)
	suite.Equal([]byte("value04"), value)
	value, err = suite.storage.Get([]byte("key05"))
	suite.Nil(err)
	suite.Equal([]byte("value05"), value)
	value, err = suite.storage.Get([]byte("key06"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)

	err = suite.storage.BeginBatch()
	suite.Nil(err)

	err = suite.storage.Delete([]byte("key04"))
	suite.Nil(err)
	err = suite.storage.Put([]byte("key06"), []byte("value06"))
	suite.Nil(err)

	err = suite.storage.Close()
	suite.Nil(err)
	storage, err := NewStorage(DBPATH, suite.t)
	suite.Require().Nil(err)
	suite.storage = storage

	err = suite.storage.CommitBatch()
	suite.NotNil(err)

	value, err = suite.storage.Get([]byte("key04"))
	suite.Nil(err)
	suite.Equal([]byte("value04"), value)
	value, err = suite.storage.Get([]byte("key05"))
	suite.Nil(err)
	suite.Equal([]byte("value05"), value)
	value, err = suite.storage.Get([]byte("key06"))
	suite.Nil(err)
	suite.Equal([]byte{}, value)
}

func (suite *StorageTestSuite) TearDownTest() {
	err := suite.storage.Close()
	suite.Nil(err)
	cmd := exec.Command("rm", "-r", DBPATH)
	err = cmd.Run()
	suite.Require().Nil(err)
}

func TestStorageTestSuite(t *testing.T) {
	suite.Run(t, &StorageTestSuite{t: LevelDBStorage})
	suite.Run(t, &StorageTestSuite{t: RocksDBStorage})
}

func BenchmarkStorage(b *testing.B) {
	for _, t := range []StorageType{LevelDBStorage, RocksDBStorage} {
		storage, err := NewStorage(DBPATH, t)
		if err != nil {
			b.Fatalf("Failed to new storage: %v", err)
		}

		keys := make([][]byte, b.N)
		values := make([][]byte, b.N)
		for i := 0; i < 1000000; i++ {
			key := make([]byte, 32)
			value := make([]byte, 32)
			rand.Read(key)
			rand.Read(value)
			keys = append(keys, key)
			values = append(values, value)
		}

		b.Run(reflect.TypeOf(storage.StorageBackend).String()+"Put", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				storage.Put(keys[i], values[i])
			}
		})
		b.Run(reflect.TypeOf(storage.StorageBackend).String()+"Get", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				storage.Get(keys[i])
			}
		})
		b.Run(reflect.TypeOf(storage.StorageBackend).String()+"Delete", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				storage.Delete(keys[i])
			}
		})

		storage.Close()
		cmd := exec.Command("rm", "-r", DBPATH)
		cmd.Run()
	}
}
