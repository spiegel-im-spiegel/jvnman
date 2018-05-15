package facade

import (
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

const (
	dbpath = "./jvndb.sqlite3"
)

var (
	//Name is applicatin name
	Name = "jvnman"
	//Version is version for applicatin
	Version = "v0.2.0"
)

var (
	cfgFile string //config file
	//cui     = rwi.New() //CUI instance
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "JVN database management",
		Long:  "JVN database management",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.jvnman.yaml)")
	rootCmd.PersistentFlags().StringP("dbfile", "", dbpath, "database file path")
	rootCmd.PersistentFlags().StringP("logfile", "", "", "logfile path (default standard error)")
	rootCmd.PersistentFlags().StringP("loglevel", "", "error", "log level: trace/debug/info/warn/error/fatal")
	viper.BindPFlag("dbfile", rootCmd.PersistentFlags().Lookup("dbfile"))
	viper.BindPFlag("logfile", rootCmd.PersistentFlags().Lookup("logfile"))
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
	rootCmd.AddCommand(newVersionCmd(ui))
	rootCmd.AddCommand(newInitCmd(ui))
	rootCmd.AddCommand(newUpdateCmd(ui))
	rootCmd.AddCommand(newListCmd(ui))
	rootCmd.AddCommand(newInfoCmd(ui))

	return rootCmd
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
			panic(err)
		}
		// Search config in home directory with name ".jvnman.yaml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".jvnman")
	}
	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadInConfig() // If a config file is found, read it in.
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
}

/* Copyright 2018 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
