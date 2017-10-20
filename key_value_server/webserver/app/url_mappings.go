package app

import (
	"github.com/javiroberts/key_value_server/webserver/controllers/identifications"
	"github.com/javiroberts/key_value_server/webserver/controllers/ping"
)

func ConfigureMappings() {
	Router.GET("/ping", ping.Ping)
	Router.POST("/identifications", identifications.HandleIdentification)
	Router.GET("/identifications/:idType/:idNumber", identifications.GetIdentification)
}
