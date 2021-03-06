package boltdb

import (
	"os"
	"testing"

	"github.com/rodrigodiez/zorro/lib/random"
)

func BenchmarkBoltDb(b *testing.B) {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := New(path)
	defer storage.Close()

	for i := 0; i < b.N; i++ {
		key := random.NewString(24)
		value := random.NewString(24)

		storage.LoadOrStore(key, value)
	}
}

func BenchmarkBoltDBParallel(b *testing.B) {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := New(path)
	defer storage.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := random.NewString(24)
			value := random.NewString(24)

			storage.LoadOrStore(key, value)
		}
	})
}
