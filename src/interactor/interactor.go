package interactor

import (
	"go-todo-api/src/domain/repository"
	"go-todo-api/src/infrastructure/persistence/datastore"
	"go-todo-api/src/usecase"
	"go-todo-api/src/userinterface/http/handler"

	"github.com/jinzhu/gorm"
)

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *gorm.DB
}

func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	handler.UserHandler
	handler.HcHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	appHandler.HcHandler = handler.NewHcHandler()
	return appHandler
}

func (interactor *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(interactor.Conn)
}

func (interactor *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(interactor.NewUserRepository())
}

func (interactor *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(interactor.NewUserUseCase())
}
