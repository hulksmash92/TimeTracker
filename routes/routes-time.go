package routes

import (
	"encoding/json"
	"net/http"
	"time"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"
)

func getTimeEntries(w http.ResponseWriter, r *http.Request) {
	token, err := parseTokenFromCookie(r)
	helpers.HandleError(err)

	ct, err := github.CheckToken(token)
	helpers.HandleError(err)

	dateFrom, err := time.Parse("2006-08-25", r.URL.Query().Get("dateFrom"))
	if err != nil {
		dateFrom = time.Now().AddDate(0, 0, -29)
	}

	dateTo, err := time.Parse("2006-08-25", r.URL.Query().Get("dateTo"))
	if err != nil {
		dateTo = time.Now()
	}

	userId := db.GetUserId(*ct.User.Login)
	timeData := db.GetTimeEntries(userId, dateFrom, dateTo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeData)
}
