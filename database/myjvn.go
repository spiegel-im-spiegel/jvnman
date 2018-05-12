package database

import (
	"time"

	"github.com/spiegel-im-spiegel/go-myjvn/rss"
	"github.com/spiegel-im-spiegel/go-myjvn/rss/option"
	"github.com/spiegel-im-spiegel/go-myjvn/status"
	"github.com/spiegel-im-spiegel/go-myjvn/vuldef"
)

//GetJVNRSS returns rss.JVNRSS (datePublished condition)
func (db *DB) GetJVNRSS(start time.Time, month bool, keyword string) (*rss.JVNRSS, error) {
	opt := option.New(
		option.WithRangeDatePublicMode(option.NoRange),
		option.WithRangeDateFirstPublishedMode(option.NoRange),
	)
	if month {
		opt.SetRangeDatePublishedMode(option.RangeMonth)
	} else if !start.IsZero() {
		opt.SetRangeDatePublishedPeriod(start, time.Now())
	}
	if len(keyword) > 0 {
		opt.SetKeyword(keyword)
	}
	return db.getJVNRSS(opt)
}

//GetJVNRSSByKeyword returns rss.JVNRSS (keyword condition)
func (db *DB) GetJVNRSSByKeyword(keyword string) (*rss.JVNRSS, error) {
	return db.getJVNRSS(option.New(
		option.WithRangeDatePublicMode(option.NoRange),
		option.WithRangeDatePublishedMode(option.NoRange),
		option.WithRangeDateFirstPublishedMode(option.NoRange),
		option.WithKeyword(keyword),
	))
}

func (db *DB) getJVNRSS(opt *option.Option) (*rss.JVNRSS, error) {
	startItem := 1
	maxItem := 0
	jvnrss := (*rss.JVNRSS)(nil)
	for {
		opt.SetStartItem(startItem)
		if startItem == 1 {
			xml, err := db.GetAPI().VulnOverviewListXML(opt)
			if err != nil {
				return jvnrss, err
			}
			stat, err := status.Unmarshal(xml)
			if err != nil {
				return jvnrss, err
			}
			if err2 := stat.GetError(); err2 != nil {
				return jvnrss, err2
			}
			maxItem = stat.Status.TotalRes
			r, err := rss.Unmarshal(xml)
			if err != nil {
				return jvnrss, err
			}
			startItem += len(r.Items)
			jvnrss = r
		} else {
			r, err := db.GetAPI().VulnOverviewList(opt)
			if err != nil {
				return jvnrss, err
			}
			startItem += len(r.Items)
			jvnrss.Append(r)
		}
		if maxItem < startItem {
			break
		}
	}
	return jvnrss, nil
}

//GetVULDEF returns vuldef.VULDEF
func (db *DB) GetVULDEF(idlist []string) (*vuldef.VULDEF, error) {
	vulnInfo := (*vuldef.VULDEF)(nil)
	for i := 0; i < len(idlist); i += vuldef.MaxItems {
		var ids []string
		if i+vuldef.MaxItems < len(idlist) {
			ids = idlist[i : i+vuldef.MaxItems]
		} else {
			ids = idlist[i:]
		}
		vuln, err := db.GetAPI().VulnDetailInfo(ids)
		if err != nil {
			return vulnInfo, err
		}
		if vulnInfo == nil {
			vulnInfo = vuln
		} else {
			vulnInfo.Append(vuln)
		}
	}
	return vulnInfo, nil
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
