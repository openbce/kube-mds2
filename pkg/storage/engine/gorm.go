package engine

import "openbce.io/kmds/pkg/storage"

func init() {
	storage.Register("gorm", gormNew)
}

func gormNew(backend string) storage.Storage {
	return &Gorm{}
}

type Gorm struct {
}
