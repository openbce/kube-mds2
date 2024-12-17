package engine

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"openbce.io/kmds/pkg/storage"
)

func init() {
	storage.Register("gorm", gormNew)
}

func gormNew(backend string) storage.Storage {
	db, err := gorm.Open(sqlite.Open(backend), &gorm.Config{})
	if err != nil {
		return nil
	}

	db.AutoMigrate(&storage.Record{})

	return &Gorm{
		db: db,
	}
}

type Gorm struct {
	db *gorm.DB
}

func (g *Gorm) CreateOrUpdate(r *storage.Record) (*storage.Record, error) {
	res := g.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "lease", "create_revision", "mod_revision"}),
	}).Create(r)

	if res.Error != nil {
		return nil, res.Error
	}

	return r, nil
}
