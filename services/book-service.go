package services

import (
	"go/auth/entities"
	"go/auth/interfaces"
)

type BookService interface {
	CreateRecord(b *entities.Book) error
	UpdateById(b *entities.Book, id int) (entities.Book, error)
	DeleteById(id int) error
	GetByID(id int) (*entities.Book, error)
}

type bookService struct {
	repo interfaces.BookRepository
}

func NewBookService(repo interfaces.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) CreateRecord(b *entities.Book) error {
	return s.repo.Create(b)
}

func (s *bookService) UpdateById(b *entities.Book, id int) (entities.Book, error) {
	return s.repo.Update(b, id)
}

func (s *bookService) DeleteById(id int) error {
	return s.repo.Delete(id)
}

func (s *bookService) GetByID(id int) (*entities.Book, error) {
	return s.repo.FindByID(id)
}
