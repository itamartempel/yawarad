package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
	"golang.org/x/net/websocket"
	"io"
)

// RequestChangesController implements the request-changes resource.
type RequestChangesController struct {
	*goa.Controller
}

// NewRequestChangesController creates a request-changes controller.
func NewRequestChangesController(service *goa.Service) *RequestChangesController {
	return &RequestChangesController{Controller: service.NewController("RequestChangesController")}
}

// Subscribe runs the subscribe action.
func (c *RequestChangesController) Subscribe(ctx *app.SubscribeRequestChangesContext) error {
	c.SubscribeWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// SubscribeWSHandler establishes a websocket connection to run the subscribe action.
func (c *RequestChangesController) SubscribeWSHandler(ctx *app.SubscribeRequestChangesContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// RequestChangesController_Subscribe: start_implement

		// Put your logic here

		ws.Write([]byte("subscribe request-changes"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// RequestChangesController_Subscribe: end_implement
	}
}
