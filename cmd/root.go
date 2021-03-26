package cmd

import (
	"fmt"
	"github.com/jsdaniell/whats-cli/cmd/whats_connection"
	"github.com/jsdaniell/whats-cli/cmd/whats_messages"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var prefix string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "whats-cli",
	Short: "This is a simple cli tool to generate different boilerplate for different type of projects.",
	Long: `The original idea comes from the necessity of create different projects with different newer frameworks that's
	coming from the newer technologies.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Welcome to your CLI application!`)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.whats-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	RootCmd.AddCommand(whats_connection.WhatsConnect)
	RootCmd.AddCommand(whats_connection.WhatsConnectQR)
	RootCmd.AddCommand(whats_connection.WhatsVersion)
	RootCmd.AddCommand(whats_connection.WhatsDisconnect)

	RootCmd.AddCommand(whats_messages.WhatsSendMessage)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".whats-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".whats-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
