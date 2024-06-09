package request

type (
	CreateUserRequest struct {
		Name          string `json:"name" validate:"required,max=50"`
		Email         string `json:"email" validate:"required,email,max=100"`
		Password      string `json:"password" validate:"required,max=100"`
		ImageFilePath string `json:"image_file_path" validate:"max=100"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email,max=100"`
		Password string `json:"password" validate:"required,max=100"`
	}

	UpdateUserRequest struct {
		Name          string `json:"name" validate:"required,max=50"`
		ImageFilePath string `json:"image_file_path" validate:"max=100"`
	}
)
