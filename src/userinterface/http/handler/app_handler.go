package handler

type AppHandler interface {
	UserHandler
	HcHandler
	TaskGroupHandler
	TaskHandler
}

type appHandler struct {
	UserHandler
	HcHandler
	TaskGroupHandler
	TaskHandler
}

func NewAppHandler(
	userHandler UserHandler,
	hcHandler HcHandler,
	taskGroupHandler TaskGroupHandler,
	taskHandler TaskHandler,
) AppHandler {
	return &appHandler{userHandler, hcHandler, taskGroupHandler, taskHandler}
}
