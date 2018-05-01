package facade

import (
	"io"
	"os"
	"syscall"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/jvnman/database"
)

func getDB(cmd *cobra.Command, defWriter io.Writer, cflag bool) (*database.DB, error) {
	dbf := viper.GetString("dbfile")
	if len(dbf) == 0 {
		return nil, errors.Wrap(os.ErrInvalid, "--dbfile")
	}
	level := database.GetLogLevel(viper.GetString("loglevel"))
	logfname := viper.GetString("logfile")
	logf := defWriter
	if len(logfname) > 0 {
		file, err := rotatelogs.New(logfname)
		if err != nil {
			return nil, err
		}
		logf = file
	}
	logger := database.NewLogger(logf, level)
	if cflag {
		logger.Println("Remove", dbf)
		if err := os.Remove(dbf); err != nil {
			switch e := err.(type) {
			case *os.PathError:
				if errno, ok := e.Err.(syscall.Errno); ok {
					if errno != syscall.ENOENT {
						logger.Fatalln(err)
						return nil, err
					}
				} else {
					logger.Fatalln(err)
					return nil, err
				}
			default:
				logger.Fatalln(err)
				return nil, err
			}
		}
	}
	return database.New(dbf, logger)
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
