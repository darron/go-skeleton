// Copyright © 2018 Salesforce
// +build linux darwin freebsd

package cmd

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"

	_ "github.com/heroku/x/hmetrics/onload" // golint
)

var startCommand = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "Start binary-name.",
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		Start()
	},
}

func init() {
	RootCmd.AddCommand(startCommand)
}

// Start starts up misr.
func Start() {

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out},"forwarded_for":"${header:X-Forwarded-For}"}` + "\n",
	}))

	portNum := os.Getenv("PORT")

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "binary-name")
	})

	e.GET("/example", service.GetExample)

	e.Logger.Fatal(e.Start(":" + portNum))
}
