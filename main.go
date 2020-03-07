// Copyright 2020 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var VERSION = ""   // Use git tag: $(git describe --always)
var GITCOMMIT = "" // Use current git commit id: $(git rev-parse --short HEAD)

func main() {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version prints the httpserver version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("httpserver version: " + VERSION + " (" + GITCOMMIT + ")")
			os.Exit(0)
		},
	}

	var rootCmd = &cobra.Command{
		Use: "httpserver",
		Run: runCmd,
	}

	rootCmd.AddCommand(versionCmd)

	rootCmd.Flags().StringP("port", "p", "", "listen port (default: 8080 [http] & 4430 [https])")
	rootCmd.Flags().StringP("path", "d", ".", "directory of files")
	rootCmd.Flags().BoolP("https", "s", false, "enable https")
	rootCmd.Flags().BoolP("browser", "b", false, "open site on browser")
	rootCmd.Flags().Bool("save", false, "save ssl certificate for reuse")
	rootCmd.Flags().BoolP("remove", "r", false, "remove stored ssl certificate")
	rootCmd.Flags().BoolP("quiet", "q", false, "don't produce logs")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
