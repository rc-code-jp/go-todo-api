package usecase

import (
	"context"
	"errors"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/common"
	"go-todo-api/src/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	CreateUser(
		ctx context.Context,
		name string,
		email string,
		password string,
		imageFilePath string,
	) (userID int, token string, err error)
	Login(ctx context.Context, email string, password string) (userID int, token string, err error)
	UpdateUser(ctx context.Context, m *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{repository}
}

// ユーザー作成
func (usecase *userUseCase) CreateUser(ctx context.Context, name, email, password, imageFilePath string) (userID int, token string, err error) {
	// パスワード暗号化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, "", err
	}

	user := model.User{
		Name:           name,
		Email:          email,
		HashedPassword: string(hashedPassword),
		ImageFilePath:  imageFilePath,
	}

	_, createErr := usecase.UserRepository.Create(ctx, &user)
	if createErr != nil {
		return 0, "", createErr
	}

	// JWTトークン生成
	token, err = common.CreateJwtToken(user.ID)

	return user.ID, token, err
}

// ログイン
func (usecase *userUseCase) Login(ctx context.Context, email string, password string) (userID int, token string, err error) {
	user, err := usecase.UserRepository.FetchByEmail(ctx, email)
	if err != nil {
		return 0, "", errors.New("メールアドレスまたはパスワードに誤りがあります。")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return 0, "", errors.New("メールアドレスまたはパスワードに誤りがあります。")
	}

	// JWTトークン生成
	token, err = common.CreateJwtToken(user.ID)

	return user.ID, token, err
}

// ユーザー取得
func (usecase *userUseCase) GetUser(ctx context.Context, id int) (*model.User, error) {
	user, err := usecase.UserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// ユーザー更新
func (usecase *userUseCase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return usecase.UserRepository.Update(ctx, user)
}

// ユーザー削除
func (usecase *userUseCase) DeleteUser(ctx context.Context, id int) error {
	if err := usecase.UserRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
