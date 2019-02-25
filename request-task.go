package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
	"golang.org/x/net/websocket"
	"io"
)

// RequestTaskController implements the request-task resource.
type RequestTaskController struct {
	*goa.Controller
}

// NewRequestTaskController creates a request-task controller.
func NewRequestTaskController(service *goa.Service) *RequestTaskController {
	return &RequestTaskController{Controller: service.NewController("RequestTaskController")}
}

// List runs the list action.
func (c *RequestTaskController) List(ctx *app.ListRequestTaskContext) error {
	// RequestTaskController_List: start_implement

	// Put your logic here

	res := app.BranchRequestTaskCollection{}
	return ctx.OK(res)
	// RequestTaskController_List: end_implement
}

// Log runs the log action.
func (c *RequestTaskController) Log(ctx *app.LogRequestTaskContext) error {
	// RequestTaskController_Log: start_implement

	// Put your logic here

	res := &app.BranchRequestTask{}
	return ctx.OK(res)
	// RequestTaskController_Log: end_implement
}

// Show runs the show action.
func (c *RequestTaskController) Show(ctx *app.ShowRequestTaskContext) error {
	// RequestTaskController_Show: start_implement

	// Put your logic here

	res := &app.BranchRequestTask{}
	return ctx.OK(res)
	// RequestTaskController_Show: end_implement
}

// StreamLog runs the stream-log action.
func (c *RequestTaskController) StreamLog(ctx *app.StreamLogRequestTaskContext) error {
	c.StreamLogWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// StreamLogWSHandler establishes a websocket connection to run the stream-log action.
func (c *RequestTaskController) StreamLogWSHandler(ctx *app.StreamLogRequestTaskContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// RequestTaskController_StreamLog: start_implement

		// Put your logic here

		ws.Write([]byte("stream-log request-task"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// RequestTaskController_StreamLog: end_implement
	}
}
