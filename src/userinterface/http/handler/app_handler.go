package handler

type AppHandler interface {
	UserHandler
	HcHandler
}

type appHandler struct {
	UserHandler
	HcHandler
}

func NewAppHandler(
	userHandler UserHandler,
	hcHandler HcHandler,
) AppHandler {
	return &appHandler{userHandler, hcHandler}
}
