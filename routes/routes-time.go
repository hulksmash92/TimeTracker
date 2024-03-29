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
	// Get the id of the authenticated user making the request
	userId := getUserId(r)

	// Parse the from date param, if the format is incorrect
	// or the value is nil/missing, set it to be 29 days ago
	dateFrom, err := time.Parse(dtParamLayout, r.URL.Query().Get("from"))
	if err != nil {
		dateFrom = time.Now().AddDate(0, 0, -29)
	}

	// Parse the to date param, if the format is incorrect
	// or the value is nil/missing, set it to now
	dateTo, err := time.Parse(dtParamLayout, r.URL.Query().Get("to"))
	if err != nil {
		dateTo = time.Now()
	}

	// Parse the current page index param, set to the first page
	// if the value is missing/nil
	pageIndex, err := strconv.ParseUint(r.URL.Query().Get("pageIndex"), 10, 32)
	if err != nil {
		pageIndex = 0
	}

	// Parse the current page index param, set to the 10 records per page
	// if the value missing/nil
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("pageSize"), 10, 32)
	if err != nil {
		pageSize = 10
	}

	// Parse the name of the column to sort by, default to creation date if missing/blank
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "created"
	}

	// Parse the sort descending param and default to true if the value is missing
	sortDesc, err := strconv.ParseBool(r.URL.Query().Get("sortDesc"))
	if err != nil {
		sortDesc = true
	}

	// Build the object that defines our pagination and sorting params for the database
	pagination := db.Pagination{
		PageSize:  uint(pageIndex),
		PageIndex: uint(pageSize),
		Sort:      sort,
		SortDesc:  sortDesc,
	}

	// Get our available rows and current page of rows from the database
	rowCount, timeData := db.GetTimeEntries(userId, dateFrom, dateTo, pagination)

	// Construct our response body object and send back to the caller
	// indicating that the body is encoded as JSON
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

// Handles the HTTP DELETE request for deleting a user's existing time entry
func deleteTime(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the current user
	userId := getUserId(r)

	// Get the variable in the URL i.e. {id}
	vars := mux.Vars(r)

	// Parse the time entries id from the URL and check for any errors
	entryId, err := strconv.ParseUint(vars["id"], 10, 32)
	helpers.HandleError(err)

	// Delete the time entry from the database, if successful the
	// returned error will be nil
	err = db.DeleteTimeEntry(userId, uint(entryId))

	// Check if an error is present
	helpers.HandleError(err)

	// return true in the data property of the bodyif no errors are present
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
