package apiv1

import (
	"github.com/sebdehne/ordercapture/server"
	"github.com/kataras/iris"
)

func Config() server.Api {
	return server.Api{Version: 1, Routes: routes()}
}

func routes() []server.Route {
	return []server.Route{
		server.Route{iris.MethodGet, "/test", testHandler},
	}
}

func testHandler(c *iris.Context) {
	c.Write("Test route")
}
