package facade

import (
	"os"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/jvnman/database"
)

//newVersionCmd returns cobra.Command instance for show sub-command
func newInitCmd(ui *rwi.RWI) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize JVN database",
		Long:  "Initialize JVN database",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbf := viper.GetString("dbfile")
			if len(dbf) == 0 {
				return errors.Wrap(os.ErrInvalid, "--dbfile")
			}
			if v, err := cmd.Flags().GetBool("verbose"); err != nil {
				return errors.Wrap(os.ErrInvalid, "--verbose")
			} else if v {
				ui.Outputln("Initialize", dbf)
			}
			if err := os.Remove(dbf); err != nil {
				switch e := err.(type) {
				case *os.PathError:
					if errno, ok := e.Err.(syscall.Errno); ok {
						if errno != syscall.ENOENT {
							return err
						}
					} else {
						return err
					}
				default:
					return err
				}
			}

			db, err := database.New(dbf)
			if err != nil {
				return err
			}
			defer db.Close()

			err = db.Initialize()
			return err
		},
	}
	//initCmd.PersistentFlags().StringP("dbfile", "f", dbpath, "database file name")
	//viper.BindPFlag("dbfile", initCmd.PersistentFlags().Lookup("dbfile"))

	return initCmd
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
