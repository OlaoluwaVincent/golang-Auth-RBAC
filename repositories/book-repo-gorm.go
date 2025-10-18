package repositories

import (
	"errors"
	"go/auth/entities"
	"go/auth/interfaces"

	"github.com/jinzhu/gorm"
)

type GormBookRepository struct {
	db *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) interfaces.BookRepository {
	return &GormBookRepository{db: db}
}

func (r *GormBookRepository) Create(u *entities.Book) error {
	return r.db.Create(u).Error
}

func (r *GormBookRepository) FindByID(id int) (*entities.Book, error) {
	var book entities.Book
	if err := r.db.First(&book, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &entities.Book{
		Title:  book.Title,
		Author: book.Author,
		ID:     book.ID,
	}, nil
}

func (r *GormBookRepository) Update(u *entities.Book, id int) (entities.Book, error) {
	var book entities.Book

	if err := r.db.First(&book, id).Error; err != nil {
		return book, errors.New("not found")
	}

	if u.Title != "" {
		book.Title = u.Title
	}
	if u.Author != "" {
		book.Author = u.Author
	}

	if err := r.db.Save(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *GormBookRepository) Delete(id int) error {
	var book entities.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return errors.New("not found")
	}
	return r.db.Delete(&book).Error
}
