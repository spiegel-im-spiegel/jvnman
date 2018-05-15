package report

//VulnInfo is "Vulnview" dataset for report
type VulnInfo struct {
	ID          string
	Title       string
	Description string
	URI         string
	Impact      string
	Solution    string
	CVSSVector  string
	Severity    string
	DatePublic  string
	DatePublish string
	DateUpdate  string
}

//AffectInfo is "Affected" dataset for report
type AffectInfo struct {
	ID            string
	Name          string
	ProductName   string
	VersionNumber string
}

//CVSSInfo is "CVSS" dataset for report
type CVSSInfo struct {
	ID         string
	Version    string
	BaseVector string
	BaseReport string
	BaseScore  string
	Severity   string
}

//RelatedInfo is "Related" dataset for report
type RelatedInfo struct {
	ID        string
	Type      string
	Name      string
	VulinfoID string
	Title     string
	URL       string
}

//VulnDetail is detail infot for report
type VulnDetail struct {
	Info       VulnInfo
	Affects    []AffectInfo
	CVSS       CVSSInfo
	Relattions []RelatedInfo
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
