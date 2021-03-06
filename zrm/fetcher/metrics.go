// Copyright 2015 The zerium Authors
// This file is part of the zerium library.
//
// The zerium library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The zerium library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the zerium library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/apolo-technologies/zerium/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("zrm/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("zrm/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("zrm/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("zrm/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("zrm/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("zrm/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("zrm/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("zrm/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("zrm/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("zrm/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("zrm/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("zrm/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("zrm/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("zrm/fetcher/filter/bodies/out")
)
