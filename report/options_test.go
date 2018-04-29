package report

import "testing"

func TestFormat(t *testing.T) {
	testCase := []struct {
		f Format
		s string
	}{
		{f: FormUnknown, s: "unknown"},
		{f: FormHTML, s: "html"},
		{f: FormMarkdown, s: "markdown"},
		{f: FormCSV, s: "csv"},
	}

	for _, tc := range testCase {
		f := TypeofFormat(tc.s)
		if f != tc.f {
			t.Errorf("TypeofFormat(\"%v\")  = %v, want %v.", tc.s, int(f), int(tc.f))
		} else {
			s := f.String()
			if s != tc.s {
				t.Errorf("Format.String(%v)  = \"%v\", want \"%v\".", int(f), s, tc.s)
			}
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
