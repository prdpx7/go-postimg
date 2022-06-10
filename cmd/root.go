/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"time"

	"github.com/briandowns/spinner"
	"github.com/prdpx7/go-postimg/imgur"
	"github.com/spf13/cobra"
)

func startSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "Uploading image"
	s.Start()
	return s
}

func stopSpinner(s *spinner.Spinner) {
	s.Stop()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-postimg",
	Short: "Upload images to imgur anonymously ",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("image path is missing")
			return
		}
		_spinner := startSpinner()
		link := imgur.UploadImage(args[0])
		stopSpinner(_spinner)
		fmt.Println(link)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-postimg.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
