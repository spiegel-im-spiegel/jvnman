package facade

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/jvnman/database"
	"github.com/spiegel-im-spiegel/jvnman/report"
)

//newListCmd returns cobra.Command instance for show sub-command
func newListCmd(ui *rwi.RWI) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List JVN data",
		Long:  "List JVN data",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbf := viper.GetString("dbfile")
			if len(dbf) == 0 {
				return errors.Wrap(os.ErrInvalid, "--dbfile")
			}
			debugf := ui.ErrorWriter()
			if debug, err := cmd.Flags().GetBool("debug"); err != nil {
				return errors.Wrap(err, "--debug")
			} else if !debug {
				debugf = nil
			}
			days, err := cmd.Flags().GetInt("range")
			if err != nil {
				return errors.Wrap(err, "--range")
			}
			score, err := cmd.Flags().GetFloat64("score")
			if err != nil {
				return errors.Wrap(err, "--score")
			}
			v, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				return errors.Wrap(err, "--verbose")
			}
			f, err := cmd.Flags().GetString("form")
			if err != nil {
				return errors.Wrap(err, "--form")
			}
			form := report.TypeofFormat(f)
			if form == report.FormUnknown {
				return errors.New("not support format: " + f)
			}

			db, err := database.New(dbf, debugf)
			if err != nil {
				return err
			}
			r, err := report.ListData(db, days, score, form, v)
			if err != nil {
				return err
			}
			ui.WriteFrom(r)

			return nil
		},
	}
	listCmd.Flags().IntP("range", "r", 3, "list the data for the past days")
	listCmd.Flags().Float64P("score", "s", 0.0, "minimum score of CVSS")
	listCmd.Flags().BoolP("verbose", "v", false, "verbose mode")
	listCmd.Flags().StringP("form", "f", "markdown", "output format: html/markdown/csv")

	return listCmd
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
