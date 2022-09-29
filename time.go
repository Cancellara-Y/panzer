package panzer

import (
	"strconv"
	"time"
)

func GetMilliSecsTimeDateRange(beginMilliTime, endMilliTime int64 /*unix time  milli seconds*/) (out []uint32) {

	startTime := beginMilliTime / 1000
	endTime := endMilliTime / 1000

	if startTime > endTime {
		return []uint32{}
	}
	st := time.Unix(startTime, 0)

	endDate64, _ := strconv.ParseUint(time.Unix(endTime, 0).Format("20060102"), 10, 64)
	endDate := uint32(endDate64)

	startDate := uint32(0)

	for startDate != endDate {

		startDate64, _ := strconv.ParseUint(st.Format("20060102"), 10, 64)
		startDate = uint32(startDate64)
		out = append(out, startDate)
		st = st.Add(time.Second * 3600 * 24)

		if startDate >= endDate {
			return
		}
	}
	return
}

func GetSecsTimeDateRange(startTime, endTime int64 /*unix time seconds*/) (out []uint32) {
	if startTime > endTime {
		return []uint32{}
	}
	st := time.Unix(startTime, 0)

	endDate64, _ := strconv.ParseUint(time.Unix(endTime, 0).Format("20060102"), 10, 64)
	endDate := uint32(endDate64)

	startDate := uint32(0)

	for startDate != endDate {

		startDate64, _ := strconv.ParseUint(st.Format("20060102"), 10, 64)
		startDate = uint32(startDate64)
		out = append(out, startDate)
		st = st.Add(time.Second * 3600 * 24)

		if startDate >= endDate {
			return
		}
	}
	return
}
