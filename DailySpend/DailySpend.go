package dailyspend

import "time"

type DailySpend struct {
	Date        time.Time
	SpentAmount int
	GoalAmount  int
}
