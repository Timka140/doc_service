package task_local

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (t *TTaskLocal) loadBase() error {

	db, err := gorm.Open(sqlite.Open(t.nameStore), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("NewDB(): открытие базы данных, err=%w", err)
	}
	t.db = db.Session(&gorm.Session{})

	return nil
}
