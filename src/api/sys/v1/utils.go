/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v3"
)

var start time.Time = time.Now()

const (
	UptimeRequiredPriv int = 3
)

func RegisterUtilRoutes(r *router.Router) {
	utils := r.Group("/v1alpha1/utils")

	utils.Handle(fasthttp.MethodGet, "/uptime", Uptime)
}

type UptimeSchema struct {
	Hours   float64 `json:"uptime_hours" yaml:"UptimeHours"`
	Minutes float64 `json:"uptime_minutes" yaml:"UptimeMinutes"`
	Seconds float64 `json:"uptime_seconds" yaml:"UptimeSeconds"`
	String  string  `json:"string" yaml:"String"`
}

// Gives the uptime for the service following the above uptime schema.
func Uptime(ctx *fasthttp.RequestCtx) {
	if Authenticate(ctx, UptimeRequiredPriv) {
		accept := string(ctx.Request.Header.Peek("Accept"))

		uptime := time.Now().Sub(start)

		u := &UptimeSchema{
			Hours:   uptime.Hours(),
			Minutes: uptime.Minutes(),
			Seconds: uptime.Seconds(),
			String:  uptime.String(),
		}

		switch accept {
		case "application/json":
			{
				// marshal with json.
				d, e := json.Marshal(u)
				if e != nil {
					ctx.SetStatusCode(http.StatusInternalServerError)
					return
				}

				ctx.Write(d)
				ctx.SetStatusCode(http.StatusOK)
				return
			}
		case "application/yaml":
			{
				// Marshal with yaml.
				d, e := yaml.Marshal(u)
				if e != nil {
					ctx.SetStatusCode(http.StatusInternalServerError)
					return
				}

				ctx.Write(d)
				ctx.SetStatusCode(http.StatusOK)
				return
			}
		default:
			{
				// Just write the current uptime stringified by default. That can be
				// read in with regex if needed, or they just use the JSON schema.
				ctx.WriteString(uptime.String())
				ctx.SetStatusCode(http.StatusOK)
				return
			}
		}
	}
}
