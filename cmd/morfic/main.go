/*
*	Copyright (C) 2023 Kendall Tauser
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

package main

import (
	"fmt"
	"os"

	app "github.com/fire833/morfic/cmd/morfic/app"
	"github.com/fire833/morfic/pkg"
)

var (
	version   string = "unknown"
	buildTime string = "unknown"
	commit    string = "unknown"
)

// Primary privileged bootstrap process for control plane.
func main() {

	pkg.Version = version
	pkg.BuildTime = buildTime
	pkg.Commit = commit

	if os.Getuid() != app.RootUID {
		fmt.Println("morfic control plane must be run as root.")
		os.Exit(1)
	}

	app.MorficMain()
}
