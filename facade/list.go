package facade

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/jvnman/report"
)

//newListCmd returns cobra.Command instance for show sub-command
func newListCmd(ui *rwi.RWI) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List JVN data",
		Long:  "List JVN data",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := getDB(cmd, ui.ErrorWriter(), false)
			if err != nil {
				return err
			}

			days, err := cmd.Flags().GetInt("range")
			if err != nil {
				return errors.Wrap(err, "--range")
			}
			db.GetLogger().Println("range:", days, "days")
			score, err := cmd.Flags().GetFloat64("score")
			if err != nil {
				return errors.Wrap(err, "--score")
			}
			db.GetLogger().Println("CVSS score:", score)
			p, err := cmd.Flags().GetString("product")
			if err != nil {
				return errors.Wrap(err, "--product")
			}
			db.GetLogger().Println("product:", p)
			c, err := cmd.Flags().GetString("cve")
			if err != nil {
				return errors.Wrap(err, "--cve")
			}
			db.GetLogger().Println("cve:", c)
			f, err := cmd.Flags().GetString("form")
			if err != nil {
				return errors.Wrap(err, "--form")
			}
			form := report.TypeofFormat(f)
			if form == report.FormUnknown {
				return errors.New("not support format: " + f)
			}
			db.GetLogger().Println("form:", form.String())
			v, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				return errors.Wrap(err, "--verbose")
			}
			db.GetLogger().Println("verbose:", v)

			r, err := report.ListData(db, days, score, p, c, form, v)
			if err != nil {
				db.GetLogger().Fatalln(err)
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
	listCmd.Flags().StringP("product", "p", "", "product name")
	listCmd.Flags().StringP("cve", "c", "", "CVE-ID (see https://cve.mitre.org/)")

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
