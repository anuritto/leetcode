package meetingrooms

import "sort"

type Interval struct {
	Start, End int
}

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

/**
 * @param intervals: an array of meeting time intervals
 * @return: if a person could attend all meetings
 */
func CanAttendMeetings(intervals []*Interval) bool {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[i].End
	})

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1].Start < intervals[i].End {
			return false
		}
	}
	return true
}
