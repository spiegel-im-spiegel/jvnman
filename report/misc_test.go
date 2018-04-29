package report

import "testing"

func TestGetSeverityJa(t *testing.T) {
	testCase := []struct {
		s  string
		sj string
	}{
		{s: "Critical", sj: "緊急"},
		{s: "High", sj: "重要"},
		{s: "Medium", sj: "警告"},
		{s: "Low", sj: "注意"},
		{s: "None", sj: "なし"},
		{s: "critical", sj: "緊急"},
		{s: "high", sj: "重要"},
		{s: "medium", sj: "警告"},
		{s: "low", sj: "注意"},
		{s: "none", sj: "なし"},
		{s: "unknown", sj: "unknown"},
	}

	for _, tc := range testCase {
		sj := getSeverityJa(tc.s)
		if sj != tc.sj {
			t.Errorf("getSeverityJa(\"%v\")  = \"%v\", want \"%v\".", tc.s, sj, tc.sj)
		}
	}
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
