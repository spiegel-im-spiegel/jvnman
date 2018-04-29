package report

import "strings"

var severityMap = map[string]string{
	"critical": "緊急",
	"high":     "重要",
	"medium":   "警告",
	"low":      "注意",
	"none":     "なし",
}

func getSeverityJa(s string) string {
	if sj, ok := severityMap[strings.ToLower(s)]; ok {
		return sj
	}
	return s
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
