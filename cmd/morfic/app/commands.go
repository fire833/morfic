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

package app

import (
	"fmt"
	"log"
	"os"
	"syscall"

	src "github.com/fire833/morfic/pkg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "morfic [options]",
		Short: "Run the morfic control plane.",
		Long: `This is the main binary for running the morfic linux based routing, firewalling, and secured local service control plane.
For more information about this project and for documentation, visit https://github.com/fire833/morfic or visit https://morfic.io.`,
		Version: fmt.Sprintf(": %s\nGit commit: %s\nBuild Time: %s\nGo version: %s\nOS: %s\nArchitecture: %s\nLicense: %s\n\nCopyright (C) 2022  Kendall Tauser",
			src.Version, src.Commit, src.BuildTime, src.Go, src.Os, src.Arch, src.License),
		Example: "morfic",
		Run:     rootMain,
	}

	forkNodeCmd = &cobra.Command{
		Use:    "nodefork <runtime_token>",
		Short:  "Used for trusted forking of node subprocess.",
		Run:    forkAPI,
		Hidden: true,
	}

	forkAPICmd = &cobra.Command{
		Use:    "apifork <runtime_token>",
		Short:  "Used for trusted forking of api subprocess.",
		Run:    forkNode,
		Hidden: true,
	}

	forkControllerCmd = &cobra.Command{
		Use:    "controllerfork <runtime_token>",
		Short:  "Used for trusted forking of controller subprocess.",
		Run:    forkController,
		Hidden: true,
	}
)

// The main function for the morfic control plane.
func MorficMain() {
	rootCmd.Flags().BoolVar(&src.DebugEnabled, "debug", false, "Use this subcommand to enable debugging mode for the process.")
	rootCmd.AddCommand(forkNodeCmd)
	rootCmd.AddCommand(forkAPICmd)
	rootCmd.AddCommand(forkControllerCmd)

	if e := rootCmd.Execute(); e != nil {
		log.Printf("Unable to start morfic: %v\n", e.Error())
		os.Exit(1)
	}
}

func forkAPI(cmd *cobra.Command, args []string) {

	// Compare the provided token to the root token to make sure this is a
	// legitimately spawned process.
	if args[2] == "-t" && args[3] == src.SharedToken.GetToken() {
		apiMain()
	} else {
		log.Printf("shared tokens do not match, aborting api fork\n")
		os.Exit(1)
	}

}

func forkNode(cmd *cobra.Command, args []string) {

	// Compare the provided token to the root token to make sure this is a
	// legitimately spawned process.
	if args[2] == "-t" && args[3] == src.SharedToken.GetToken() {
		nodeMain()
	} else {
		log.Printf("shared tokens do not match, aborting node fork\n")
		os.Exit(1)
	}
}

func forkController(cmd *cobra.Command, args []string) {
	// Compare the provided token to the root token to make sure this is a
	// legitimately spawned process.
	if args[2] == "-t" && args[3] == src.SharedToken.GetToken() {
		controllerMain()
	} else {
		log.Printf("shared tokens do not match, aborting controller fork\n")
		os.Exit(1)
	}
}

func rootMain(cmd *cobra.Command, args []string) {

	log.Printf("generating runtime trust tokens for bootstrapping daughter processes")
	src.GenerateRuntimeTrustAnchors()

	log.Printf("forking node process...\n")

	morfic, e := os.Executable()
	if e != nil {
		log.Printf("unable to acquire path to morfic executable: %v, exiting...\n", e)
		os.Exit(1)
	}

	// Spawn the node process first
	nodepid, nodee := syscall.ForkExec(morfic, []string{"forknode", "-t", src.SharedToken.GetToken()}, &syscall.ProcAttr{
		Sys: &syscall.SysProcAttr{
			Ptrace:     false,
			Cloneflags: syscall.CLONE_NEWIPC | syscall.CLONE_VM,
			Credential: &syscall.Credential{
				Uid:         uint32(rootUID),
				Gid:         uint32(rootGID),
				NoSetGroups: true,
			},
		},
	})

	if nodee != nil {
		log.Printf("unable to fork node process, error: %v\n", nodee.Error())
		os.Exit(1) // Kill the bootstrapping process here.
	} else {
		log.Printf("morfic node process created, PID: %d\n", nodepid)
	}

	log.Printf("forking api process...\n")

	// Spawn the api server second
	apipid, apie := syscall.ForkExec(morfic, []string{"forkapi", "-t", src.SharedToken.GetToken()}, &syscall.ProcAttr{
		Sys: &syscall.SysProcAttr{
			Ptrace:     false,
			Cloneflags: syscall.CLONE_NEWIPC | syscall.CLONE_VM,
			Credential: &syscall.Credential{
				Uid:         uint32(unprivilegedUID),
				Gid:         uint32(unprivilegedGID),
				NoSetGroups: true,
			},
		},
	})

	if apie != nil {
		log.Printf("unable to fork api process, error: %v\n", apie.Error())
		os.Exit(1) // Kill the bootstrapping process here.
	} else {
		log.Printf("morfic api process created, PID: %d\n", apipid)
	}

	log.Printf("forking controller process...\n")

	// Spawn controller manager third.
	controllerpid, controllere := syscall.ForkExec(morfic, []string{"forkcontroller", "-t", src.SharedToken.GetToken()}, &syscall.ProcAttr{
		Sys: &syscall.SysProcAttr{
			Ptrace:     false,
			Cloneflags: syscall.CLONE_NEWIPC | syscall.CLONE_VM,
			Credential: &syscall.Credential{
				Uid:         uint32(unprivilegedUID),
				Gid:         uint32(unprivilegedGID),
				NoSetGroups: true,
			},
		}})

	if controllere != nil {
		log.Printf("unable to fork controller process, error: %v\n", controllere.Error())
		os.Exit(1)
	} else {
		log.Printf("morfic controller process created, PID: %d\n", controllerpid)
	}

	log.Printf("morfic control plane bootstrapped, bootstrap process exiting")

}
