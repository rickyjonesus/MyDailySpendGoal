package profile

type PayFrequency string

const (
	Monthly   PayFrequency = "Monthly"
	Weekly                 = "Weekly"
	BiWeekly               = "BiWeekly"
	BiMonthly              = "BiMonthly"
)

type Profile struct {
	FirstName           string
	LastName            string
	EmailAddress        string
	PayFrequency        PayFrequency
	PayDays             []int
	PayPaySpendingMoney int,
	DailyGoal			decimal
	
}
