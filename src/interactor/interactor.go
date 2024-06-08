// Package interactor 簡易DIコンテナ
package interactor

import (
	"go-todo-api/src/domain/repository"
	"go-todo-api/src/infrastructure/persistence/datastore"
	"go-todo-api/src/usecase"
	"go-todo-api/src/userinterface/http/handler"

	"github.com/jinzhu/gorm"
)

// Interactor インターフェース。AppHandlerのインターフェースを保持。
type Interactor interface {
	NewAppHandler() handler.AppHandler
}

// interactor 構造体
type interactor struct {
	Conn *gorm.DB
}

// NewInteractor intractorを生成。
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

// ユーザー関連
// NewUserRepository UserRepositoryを生成。
func (interactor *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(interactor.Conn)
}

// NewUserUseCase UserUseCaseを生成。
func (interactor *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(interactor.NewUserRepository())
}

// NewUserHandler UserHandlerを生成。
func (interactor *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(interactor.NewUserUseCase())
}
