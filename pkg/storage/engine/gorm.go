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
	assignments := clause.AssignmentColumns([]string{"value", "lease"})
	assignments = append(assignments, clause.Assignments(map[string]interface{}{
		"revision":   gorm.Expr("revision + ?", 1),
		"prev_value": gorm.Expr("prev_value = value"),
	})...)

	res := g.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: assignments,
	}).Create(r)

	if res.Error != nil {
		return nil, res.Error
	}

	return r, nil
}

func (g *Gorm) Range(keys [][]byte) ([]*storage.Record, error) {
	records := []*storage.Record{}
	g.db.Where("key IN ?", keys).Find(&records)

	return records, nil
}

func (g *Gorm) Size() (int64, error) {
	// TODO (k82cn): support other db.
	sizeSQL := "SELECT SUM(pgsize) FROM dbstat"
	size := int64(-1)
	if err := g.db.Raw(sizeSQL).Row().Scan(&size); err != nil {
		return -1, err
	}

	return size, nil
}
