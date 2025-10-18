package interfaces

import "go/auth/entities"

type BookRepository interface {
	Create(book *entities.Book) error
	FindByID(id int) (*entities.Book, error)
	Update(book *entities.Book, id int) (entities.Book, error)
	Delete(id int) error
}
