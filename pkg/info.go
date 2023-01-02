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

package pkg

import (
	"fmt"
	"runtime"
)

var (
	Version   string = "unknown"         // String to pass in the version to the binary at compiletime.
	BuildTime string = "unknown"         // The build time for the control plane.
	Commit    string = "unknown"         // Git commit version of this binary.
	Go        string = runtime.Version() // Go version at runtime.
	Os        string = runtime.GOOS      // operating system for this binary
	Arch      string = runtime.GOARCH    // architecture for this binary
	License   string = "GPL V2.0"        // license for this project
)

func NewVersionString(componentName string) string {
	return fmt.Sprintf("%s - Version: %s\nCommit: %s\nGo Build: %s\nBuild Date: %sOS: %s\nArchitecture: %s\nLicense: %s\n\n",
		componentName,
		Version,
		Commit,
		Go,
		BuildTime,
		Os,
		Arch,
		License,
	)
}
