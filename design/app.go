package design

import (
	// . "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("yawarad", func() {
	Title("YAWARAD API for the lazy DBA")
	Description("Set of utiles for creating and managing databases")
	Host("localhost:8081")
	BasePath("/api")
})
