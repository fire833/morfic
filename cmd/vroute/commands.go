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

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fire833/vroute/src"
	"github.com/spf13/cobra"
)

var (
	version string = "unknown"         // String to pass in the version to the binary at compiletime.
	commit  string = "unknown"         // Git commit version of this binary.
	g       string = runtime.Version() // Go version at runtime.
	o       string = runtime.GOOS      // operating system for this binary
	arch    string = runtime.GOARCH    // architecture for this binary
	license string = "GPL V2.0"        // license for this project

	rootCmd = &cobra.Command{
		Use:   "vroute [options]",
		Short: "Run the vroute control plane.",
		Long: `This is the main binary for running the vRoute linux based routing, firewalling, and secured local service control plane.
For more information about this project and for documentation, visit https://github.com/fire833/vroute or visit https://vroute.io.`,
		Version: fmt.Sprintf(": %s\nGit commit: %s\nGo version: %s\nOS: %s\nArchitecture: %s\nLicense: %s\n\nCopyright (C) 2021  Kendall Tauser", version, commit, g, o, arch, license),
		Example: "vroute",
	}
)

// The main function for the vroute control plane.
func vrouteMain() {
	rootCmd.Flags().BoolVar(&src.DebugEnabled, "debug", false, "Use this subcommand to enable debugging mode for the process.")

	if e := rootCmd.Execute(); e != nil {
		fmt.Println("Unable to start vroute: " + e.Error())
		os.Exit(1)
	}
}
