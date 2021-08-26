package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"timetracker/db"
	"timetracker/helpers"
	"timetracker/models"

	"github.com/gorilla/mux"
)

// Handles the requests for `/api/time` and handle its
// with the correct function based on the HTTP method
func timeRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTimeEntries(w, r)
	case http.MethodPost:
		createTime(w, r)
	case http.MethodPatch:
		updateTime(w, r)
	case http.MethodDelete:
		deleteTime(w, r)
	}
}

// Handles the GET request for getting a list of time entries for the user
func getTimeEntries(w http.ResponseWriter, r *http.Request) {
	userId := getUserId(r)

	dateFrom, err := time.Parse(dtParamLayout, r.URL.Query().Get("from"))
	if err != nil {
		dateFrom = time.Now().AddDate(0, 0, -29)
	}
	dateTo, err := time.Parse(dtParamLayout, r.URL.Query().Get("to"))
	if err != nil {
		dateTo = time.Now()
	}

	pageIndex, err := strconv.ParseUint(r.URL.Query().Get("pageIndex"), 10, 32)
	if err != nil {
		pageIndex = 0
	}
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("pageSize"), 10, 32)
	if err != nil {
		pageSize = 10
	}
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "created"
	}
	sortDesc, err := strconv.ParseBool(r.URL.Query().Get("sortDesc"))
	if err != nil {
		sortDesc = true
	}

	pagination := db.Pagination{
		PageSize:  uint(pageIndex),
		PageIndex: uint(pageSize),
		Sort:      sort,
		SortDesc:  sortDesc,
	}

	rowCount, timeData := db.GetTimeEntries(userId, dateFrom, dateTo, pagination)
	res := map[string]interface{}{
		"data": map[string]interface{}{
			"rowCount": rowCount,
			"page":     timeData,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Handles the POST request for creating a new time entry
func createTime(w http.ResponseWriter, r *http.Request) {
	userId := getUserId(r)
	body := readBody(r)
	var fmtBody models.TimeEntry
	err := json.Unmarshal(body, &fmtBody)
	helpers.HandleError(err)

	fmtBody.User = models.OwnerTrimmed{Id: userId, Name: ""}
	entryId := db.CreateTimeEntry(fmtBody)
	timeData := db.GetTimeEntry(entryId)
	res := map[string]interface{}{"data": timeData}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Handles the PATCH request for updating an existing time entry
func updateTime(w http.ResponseWriter, r *http.Request) {
	userId := getUserId(r)
	vars := mux.Vars(r)
	entryId, err := strconv.ParseUint(vars["id"], 10, 32)
	helpers.HandleError(err)

	body := readBody(r)
	var fmtBody db.UpdatedTimeEntry
	err = json.Unmarshal(body, &fmtBody)
	helpers.HandleError(err)

	err = db.UpdateTimeEntry(userId, uint(entryId), fmtBody)
	helpers.HandleError(err)

	timeData := db.GetTimeEntry(uint(entryId))
	res := map[string]interface{}{"data": timeData}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Handles the DELETE request for deleting an existing time entry
func deleteTime(w http.ResponseWriter, r *http.Request) {
	userId := getUserId(r)
	vars := mux.Vars(r)
	entryId, err := strconv.ParseUint(vars["id"], 10, 32)
	helpers.HandleError(err)

	err = db.DeleteTimeEntry(userId, uint(entryId))
	helpers.HandleError(err)

	res := map[string]bool{"data": err != nil}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Handles the GET request for retreiving all tags
func getTags(w http.ResponseWriter, r *http.Request) {
	userId := getUserId(r)
	tags := db.GetTags(userId)
	res := map[string]interface{}{"data": tags}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
