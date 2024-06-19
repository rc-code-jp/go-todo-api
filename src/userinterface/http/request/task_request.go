package request

type (
	CreateTaskRequest struct {
		Title          string `json:"title" validate:"required,max=250"`
		TaskGroupId    int `json:"taskGroupId" validate:"required,min=1"`
		Date          string `json:"date" validate:""`
		Time          string `json:"time" validate:""`
	}

	UpdateTaskRequest struct {
		ID            int `json:"id" validate:"min=1"`
		TaskGroupId    int `json:"taskGroupId"`
		Title          string `json:"title" validate:"max=250"`
		Date          string `json:"date" validate:""`
		Time          string `json:"time" validate:""`
	}
)
