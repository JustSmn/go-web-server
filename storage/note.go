/*


	This code is working with the database
	There is also a data model (Note struct)


*/

package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	Id       uint16 `gorm:"primaryKey"`
	Title    string
	FullText string
}

var db *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.AutoMigrate(&Note{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	fmt.Println("Successfully connected to database!")
	return db, nil
}

func GetAllNotes(db *gorm.DB) ([]Note, error) {
	var notes []Note
	if err := db.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func CreateNote(note *Note) error {
	return db.Create(note).Error
}

func DeleteNote(id string) error {
	return db.Delete(&Note{}, id).Error
}

func GetNoteByID(id string) (Note, error) {
	var note Note
	err := db.First(&note, id).Error
	return note, err
}
