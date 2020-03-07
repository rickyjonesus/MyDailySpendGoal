package dailyspend

import (
	"encoding/json"
	"net/http"
	"time"
)

func Get(w http.ResponseWriter, r *http.Request) {

	var ret = DailySpend{
		Date:        time.Now(),
		SpentAmount: 10,
		GoalAmount:  50,
	}
	json.NewEncoder(w).Encode(ret)
}
