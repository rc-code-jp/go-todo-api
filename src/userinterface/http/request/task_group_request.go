package request

type (
	CreateTaskGroupRequest struct {
		Name          string `json:"name" validate:"required,max=100"`
	}

	UpdateTaskGroupRequest struct {
		ID            int `validate:"min=1"`
		Name          string `json:"name" validate:"max=100"`
	}
)
