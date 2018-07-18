// Copyright © 2018 Salesforce
// +build linux darwin freebsd

package service

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetExample(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, "Example values", "  ")
}
