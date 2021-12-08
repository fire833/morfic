/*
*	Copyright (C) 2021  Kendall Tauser
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

package api

import (
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var uptime time.Time = time.Now()

const (
	UptimeRequiredPriv int = 3
)

func RegisterUtilRoutes(r *router.Router) {
	utils := r.Group("/v1/utils")

	utils.Handle(fasthttp.MethodGet, "/uptime", Uptime)
}

func Uptime(ctx *fasthttp.RequestCtx) {
	if Authenticate(ctx, UptimeRequiredPriv) {

	}
}
