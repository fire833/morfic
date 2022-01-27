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

package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"

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
		Version: fmt.Sprintf(": %s\nGit commit: %s\nGo version: %s\nOS: %s\nArchitecture: %s\nLicense: %s\n\nCopyright (C) 2022  Kendall Tauser", version, commit, g, o, arch, license),
		Example: "vroute",
		Run:     rootMain,
	}

	forkNodeCmd = &cobra.Command{
		Use:   "vroute nodefork <runtime_token>",
		Short: "Used for trusted forking of node subprocess.",
		Run:   forkAPI,
	}

	forkAPICmd = &cobra.Command{
		Use:   "vroute apifork <runtime_token>",
		Short: "Used for trusted forking of api subprocess.",
		Run:   forkNode,
	}
)

// The main function for the vroute control plane.
func vrouteMain() {
	rootCmd.Flags().BoolVar(&src.DebugEnabled, "debug", false, "Use this subcommand to enable debugging mode for the process.")
	rootCmd.AddCommand(forkNodeCmd)
	rootCmd.AddCommand(forkAPICmd)

	if e := rootCmd.Execute(); e != nil {
		fmt.Println("Unable to start vroute: " + e.Error())
		os.Exit(1)
	}
}

func forkAPI(cmd *cobra.Command, args []string) {
	if args[0] == "-t" && args[1] == src.SharedToken.GetToken() {
		api_main()
	}

}

func forkNode(cmd *cobra.Command, args []string) {
	if args[0] == "-t" && args[1] == src.SharedToken.GetToken() {
		node_main()
	}
}

func rootMain(cmd *cobra.Command, args []string) {

	src.GenerateRuntimeTrustToken()

	// Spawn the node process first to start
	syscall.ForkExec("vroute", []string{"forkapi", "-t", src.SharedToken.GetToken()}, &syscall.ProcAttr{
		Sys: &syscall.SysProcAttr{
			Ptrace:     false,
			Cloneflags: syscall.CLONE_NEWIPC,
			Credential: &syscall.Credential{},
		},
	})

	syscall.ForkExec("vroute", []string{"forknode", "-t", src.SharedToken.GetToken()}, &syscall.ProcAttr{
		Sys: &syscall.SysProcAttr{
			Ptrace:     false,
			Cloneflags: syscall.CLONE_NEWIPC,
			Credential: &syscall.Credential{},
		},
	})

}
