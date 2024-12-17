package storage

type Record struct {
	ID             uint64
	CreateRevision int64
	ModRevision    int64
	Lease          int64
	Key            []byte `gorm:"index:key_index,unique"`
	Value          []byte
}

type Storage interface {
	CreateOrUpdate(r *Record) (*Record, error)
	// Delete(id uint64) (*Record, error)
	// Update(r *Record) (*Record, error)
	// Get(id uint64) (*Record, error)
}
