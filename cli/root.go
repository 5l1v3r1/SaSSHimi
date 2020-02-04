// Copyright © 2018 Raul Sampedro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"
	"github.com/op/go-logging"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verboseLevel int
var bindAddress string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Generate SSH Dynamic Tunnels when AllowTcpForwarding is off",
	Long: `This tool aims to create a Dynamic Tunnel trougth a shell channel 
of SSH using stdin and stdout to transmiti information`,
}

// Execute adds all child cli to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.SaSSHimi.yaml)")
	rootCmd.PersistentFlags().CountVarP(&verboseLevel, "verbose", "v", "verbose level")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".ssh-tunnel" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".SaSSHimi")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadInConfig()

	if verboseLevel == 0 {
		logging.SetLevel(logging.NOTICE, "SaSSHimi")
	} else if verboseLevel == 1 {
		logging.SetLevel(logging.INFO, "SaSSHimi")
	} else {
		logging.SetLevel(logging.DEBUG, "SaSSHimi")
	}
}
