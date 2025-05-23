// Copyright (C) INFINI Labs & INFINI LIMITED.
//
// The INFINI Framework is offered under the GNU Affero General Public License v3.0
// and as commercial software.
//
// For commercial licensing, contact us at:
//   - Website: infinilabs.com
//   - Email: hello@infini.ltd
//
// Open Source licensed under AGPL V3:
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

/* Copyright © INFINI Ltd. All rights reserved.
 * web: https://infinilabs.com
 * mail: hello#infini.ltd */

package elastic

import (
	"fmt"
	"infini.sh/framework/core/util"
	"testing"
)

func TestBuildSearchTermAggregations(t *testing.T) {
	aggs := BuildSearchTermAggregations([]SearchAggParam{
		{Field: "name", TermsAggParams: util.MapStr{
			"size": 100,
		}}, {
			Field: "labels.health_status",
			TermsAggParams: util.MapStr{
				"size": 10,
			},
		},
	})
	fmt.Println(aggs)
}

func TestBuildSearchTermFilter(t *testing.T) {
	filter := BuildSearchTermFilter(map[string][]string{
		"version": {"5.6.8", "2.4.6"},
		"tags":    {"v5", "infini"},
	})
	fmt.Println(filter)
}
