// Package request リクエストを表す構造体を定義
package request

type (
	// CreateUserRequest ユーザー登録リクエスト
	CreateUserRequest struct {
		Name          string `json:"name" validate:"required,max=50"`
		Email         string `json:"email" validate:"required,email,max=100"`
		Password      string `json:"password" validate:"required,max=100"`
		ImageFilePath string `json:"image_file_path" validate:"max=100"`
	}

	// LoginRequest ログインリクエスト
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email,max=100"`
		Password string `json:"password" validate:"required,max=100"`
	}

	// UpdateUserRequest ユーザー更新リクエスト
	UpdateUserRequest struct {
		ID            int    `json:"id" validate:"required,min=1"`
		Name          string `json:"name" validate:"required,max=50"`
		ImageFilePath string `json:"image_file_path" validate:"max=100"`
	}
)
