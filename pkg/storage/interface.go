package storage

type Record struct {
	ID       uint64
	Revision uint64
	Lease    uint64
	Key      []byte
	Value    []byte
}

type Storage interface {
	Create(r *Record) (*Record, error)
	// Delete(id uint64) (*Record, error)
	// Update(r *Record) (*Record, error)
	// Get(id uint64) (*Record, error)
}
