package database

import "database/sql"

var (
	//NullText is null text data
	NullText = sql.NullString{String: "", Valid: false}
	//NullInteger is null integer data
	NullInteger = sql.NullInt64{Int64: 0, Valid: false}
	//NullNumeric is null numeric data
	NullNumeric = sql.NullFloat64{Float64: 0.0, Valid: false}
)

//Text returns sql.NullString (not NULL)
func Text(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

//Integer returns sql.NullInt64 (not NULL)
func Integer(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

//Numeric returns sql.NullFloat64 (not NULL)
func Numeric(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: f, Valid: true}
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
