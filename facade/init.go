package facade

import (
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newVersionCmd returns cobra.Command instance for show sub-command
func newInitCmd(ui *rwi.RWI) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize vulnerability database",
		Long:  "Initialize vulnerability database",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := getDB(cmd, ui.ErrorWriter(), true)
			if err != nil {
				return err
			}
			if err := db.Initialize(); err != nil {
				db.GetLogger().Fatalln(err)
				return err
			}
			return nil
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
