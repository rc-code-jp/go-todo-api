package handler

type AppHandler interface {
	UserHandler
	HcHandler
	TaskGroupHandler
}

type appHandler struct {
	UserHandler
	HcHandler
	TaskGroupHandler
}

func NewAppHandler(
	userHandler UserHandler,
	hcHandler HcHandler,
	taskGroupHandler TaskGroupHandler,
) AppHandler {
	return &appHandler{userHandler, hcHandler, taskGroupHandler}
}
