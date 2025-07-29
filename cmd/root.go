/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aft",
	Short: "AWS Control Tower Account Factory for Terraform (AFT) CLI tool",
	Long: `AFT CLI is a command-line tool for managing AWS Control Tower Account Factory for Terraform.
It provides efficient management of AFT configurations, account provisioning, pipeline monitoring,
and customization management for cloud engineers and DevOps teams.

Examples:
  aft config show          Show current AFT configuration
  aft account list         List all managed accounts
  aft pipeline status      Show pipeline status
  aft dashboard            Display AFT management dashboard`,
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
	cobra.OnInitialize(initConfig)

	// Global persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aft.yaml)")
	rootCmd.PersistentFlags().StringP("profile", "p", "", "AWS profile to use")
	rootCmd.PersistentFlags().StringP("region", "r", "", "AWS region to use")
	rootCmd.PersistentFlags().StringP("output", "o", "table", "output format (table, json, yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".aft" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".aft")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
