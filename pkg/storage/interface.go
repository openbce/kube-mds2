package storage

type Record struct {
	ID        uint64
	Revision  int64
	Lease     int64
	Key       []byte `gorm:"index:key_index,unique"`
	Value     []byte
	PrevValue []byte
}

type Event struct {
	Key      []byte
	Value    []byte
	Revision int64
}

type Storage interface {
	CreateOrUpdate(r *Record) (*Record, error)
	Range(keys [][]byte) ([]*Record, error)
	Size() (int64, error)
	// Delete(id uint64) (*Record, error)
	// Update(r *Record) (*Record, error)
	// Get(id uint64) (*Record, error)
}
