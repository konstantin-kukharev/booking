package pkg

import "time"

//  import "time"

// //DaysBetween calculate days between two time.Time & return []
// func DaysBetween(from time.Time, to time.Time) []Day {
// 	days := make([]Day, 0)

// 	for d := ToDay(from); !d.After(ToDay(to)); d = d.AddDate(0, 0, 1) {
// 		days = append(days, d)
// 	}

// 	return days
// }

// func ToDay(timestamp time.Time) time.Time {
// 	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
// }

// Date return truncated time.Time in UTC Locale
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
