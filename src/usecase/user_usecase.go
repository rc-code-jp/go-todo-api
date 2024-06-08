// Package usecase Application Service層。
package usecase

import (
	"context"
	"errors"
	"os"
	"time"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/domain/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase インターフェース
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
	UpdateUser(ctx context.Context, u *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

// userUseCase 構造体
type userUseCase struct {
	repository.UserRepository
}

// NewUserUseCase UserUseCaseを生成。
func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{repository}
}

// CreateUser 登録
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
	token, err = createToken(&user)

	return user.ID, token, err
}

// Login ログイン
func (usecase *userUseCase) Login(ctx context.Context, email string, password string) (userID int, token string, err error) {
	user, err := usecase.UserRepository.FetchByEmail(ctx, email)
	if err != nil {
		return 0, "", errors.New("メールアドレスまたはパスワードに誤りがあります。")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return 0, "", errors.New("メールアドレスまたはパスワードに誤りがあります。")
	}

	// JWTトークン生成
	token, err = createToken(user)

	return user.ID, token, err
}

// createToken JWTトークンを生成
func createToken(user *model.User) (string, error) {
	// 鍵となる文字列
	secret := os.Getenv("JWT_SIGNING_KEY")

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))

	return t, err
}

// GetUser 詳細取得
func (usecase *userUseCase) GetUser(ctx context.Context, id int) (*model.User, error) {
	user, err := usecase.UserRepository.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser 更新
func (usecase *userUseCase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return usecase.UserRepository.Update(ctx, user)
}

// DeleteUser 削除
func (usecase *userUseCase) DeleteUser(ctx context.Context, id int) error {
	if err := usecase.UserRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
