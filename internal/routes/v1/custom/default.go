// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package custom

import (
	"fmt"
	"os"

	"gitea.ysicing.net/cloud/pangu/internal/routes"
	"github.com/ergoapi/util/exgin"
	"github.com/gin-gonic/gin"
)

type Custom struct{}

func init() {
	routes.Migrate(NewHandler)
}

func NewHandler() (routes.RouteRegister, error) {
	return &Custom{}, nil
}

func (api Custom) ApplyRoute(r *gin.Engine) {
	c := r.Group("")
	c.GET("/cdn-cgi/trace", api.Get)
	c.GET("/", api.Get)
}

func (api Custom) Name() string {
	return "custom"
}

// Get
// @Summary CF CDN
// @Tags custom
// @Accept application/json
// @Success 200
// @Router /cdn-cgi/trace [get]
func (api Custom) Get(c *gin.Context) {
	hostname, _ := os.Hostname()
	_, _ = fmt.Fprintln(c.Writer, "Hostname:", hostname)
	_, _ = fmt.Fprintln(c.Writer, "IP:", exgin.RealIP(c))
	_, _ = fmt.Fprintln(c.Writer, "Headers:", c.Request.Header)
	_, _ = fmt.Fprintln(c.Writer, "URL:", c.Request.URL.RequestURI())
	_, _ = fmt.Fprintln(c.Writer, "Host:", c.Request.Host)
	_, _ = fmt.Fprintln(c.Writer, "Method:", c.Request.Method)
	_, _ = fmt.Fprintln(c.Writer, "RemoteAddr:", c.Request.RemoteAddr)
	c.Request.Write(c.Writer)
}
