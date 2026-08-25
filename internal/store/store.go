package store

import (
	"coldchain/internal/domain"
	"encoding/json"
	"go.etcd.io/bbolt"
	"path/filepath"
)

var buckets = [][]byte{[]byte("parks"), []byte("vehicles"), []byte("dispatches"), []byte("events")}

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error { return s.db.Close() }
func put[T any](s *Store, b []byte, id string, v T) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(id), raw) })
}
func getAll[T any](s *Store, b []byte) ([]T, error) {
	out := []T{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(b).ForEach(func(_, v []byte) error {
			var x T
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func (s *Store) PutPark(v domain.Park) error              { return put(s, buckets[0], v.ID, v) }
func (s *Store) PutVehicle(v domain.Vehicle) error        { return put(s, buckets[1], v.ID, v) }
func (s *Store) PutDispatch(v domain.DispatchOrder) error { return put(s, buckets[2], v.ID, v) }
func (s *Store) PutEvent(v domain.DispatchEvent) error    { return put(s, buckets[3], v.ID, v) }
func (s *Store) Parks() ([]domain.Park, error)            { return getAll[domain.Park](s, buckets[0]) }
func (s *Store) Vehicles() ([]domain.Vehicle, error)      { return getAll[domain.Vehicle](s, buckets[1]) }
func (s *Store) Dispatches() ([]domain.DispatchOrder, error) {
	return getAll[domain.DispatchOrder](s, buckets[2])
}
func (s *Store) Events() ([]domain.DispatchEvent, error) {
	return getAll[domain.DispatchEvent](s, buckets[3])
}
