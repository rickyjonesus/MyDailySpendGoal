package dailyspend

import (
	"encoding/json"
	"net/http"
	"time"
)

func Get(w http.ResponseWriter, r *http.Request) {

	var emailAddress = r.URL.Query().Get("emailAddress")
	var date = r.URL.Query().Get("Date")

	var profile = profile.GetOne(emailAddress);
	
	var ret = DailySpend{
		Date:        date,
		SpentAmount: 10, //sum of transactions for day
		GoalAmount:  profile.DailyGoal, //goal from profile
	}
	json.NewEncoder(w).Encode(ret)
}
