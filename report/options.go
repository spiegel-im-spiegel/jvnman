package report

import "strings"

//Format is type of output-format
type Format int

//Output format
const (
	FormUnknown Format = iota
	FormHTML
	FormMarkdown
	FormCSV
)

var formatMap = map[string]Format{
	"html":     FormHTML,
	"markdown": FormMarkdown,
	"csv":      FormCSV,
}

//TypeofFormat returns type of Format
func TypeofFormat(s string) Format {
	if f, ok := formatMap[strings.ToLower(s)]; ok {
		return f
	}
	return FormUnknown
}

func (f Format) String() string {
	for key, value := range formatMap {
		if value == f {
			return key
		}
	}
	return "unknown"
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
