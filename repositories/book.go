package repositories

import (
	"book-crud/domain"
	"book-crud/model"
	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}
func (repo *bookRepo) GetBooks(bookID uint) []model.Book {
	var Book []model.Book
	var err error

	if bookID != 0 {
		err = repo.db.Where("id = ?", bookID).Find(&Book).Error
	} else {
		err = repo.db.Find(&Book).Error
	}

	if err != nil {
		return []model.Book{}
	}
	return Book
}
func (repo *bookRepo) CreateBook(book *model.Book) error {
	if err := repo.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *bookRepo) UpdateBook(book *model.Book) error {
	if err := repo.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *bookRepo) DeleteBook(bookID uint) error {
	var Book model.Book
	if err := repo.db.
		Where("id = ?", bookID).Delete(&Book).
		Error; err != nil {
		return err
	}
	return nil
}
