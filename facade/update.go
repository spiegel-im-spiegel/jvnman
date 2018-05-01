package facade

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newUpdateCmd returns cobra.Command instance for show sub-command
func newUpdateCmd(ui *rwi.RWI) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update JVN database",
		Long:  "Update JVN database",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := cmd.Flags().GetBool("month")
			if err != nil {
				return errors.Wrap(err, "--month")
			}

			db, err := getDB(cmd, ui.ErrorWriter(), false)
			if err != nil {
				return err
			}
			ids, err := db.Update(m)
			if err != nil {
				db.GetLogger().Fatalln(err)
				return err
			}
			if err := db.UpdateDetail(ids); err != nil {
				db.GetLogger().Fatalln(err)
				return err
			}
			return nil
		},
	}
	updateCmd.Flags().BoolP("month", "m", false, "get the data for the past month")

	return updateCmd
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
