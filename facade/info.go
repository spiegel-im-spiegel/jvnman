package facade

import (
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/jvnman/report"
)

//newInfoCmd returns cobra.Command instance for show sub-command
func newInfoCmd(ui *rwi.RWI) *cobra.Command {
	infoCmd := &cobra.Command{
		Use:   "info [flags] <JVN Vulnerability ID>",
		Short: "Output vulnerability information",
		Long:  "Output vulnerability information",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := getDB(cmd, ui.ErrorWriter(), false)
			if err != nil {
				return err
			}

			if len(args) == 0 {
				return errors.Wrap(os.ErrInvalid, "No JVN Database ID")
			} else if len(args) > 1 {
				return errors.Wrap(os.ErrInvalid, strings.Join(args, " "))
			}
			id := args[0]
			db.GetLogger().Println("JVN Database ID:", id)

			f, err := cmd.Flags().GetString("form")
			if err != nil {
				return errors.Wrap(err, "--form")
			}
			form := report.TypeofFormat(f)
			if form != report.FormHTML && form != report.FormMarkdown {
				return errors.New("not support format: " + f)
			}
			db.GetLogger().Println("form option:", form.String())
			tf, err := cmd.Flags().GetString("template")
			if err != nil {
				return errors.Wrap(err, "--template")
			}
			if len(tf) > 0 {
				db.GetLogger().Println("template option:", tf)
			}

			r, err := report.Info(db, id, tf, form)
			if err != nil {
				db.GetLogger().Fatalln(err)
				return err
			}
			ui.WriteFrom(r)

			return nil
		},
	}
	infoCmd.Flags().StringP("form", "f", "markdown", "output format: html/markdown")
	infoCmd.Flags().StringP("template", "t", "", "template file path")

	return infoCmd
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
