package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func textIssueHandler(issueGetter IssueGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		issues, err := issueGetter.GetIssues(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sort.Slice(issues, func(i, j int) bool {
			return issues[i].Importance > issues[j].Importance
		})

		for i, issue := range issues {
			fmt.Fprintf(w, "%d. %s (%d)\n", i+1, issue.Title, issue.Importance)
		}
	})
}

func jsonIssueHandler(issueGetter IssueGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		issues, err := issueGetter.GetIssues(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sort.Slice(issues, func(i, j int) bool {
			return issues[i].Importance > issues[j].Importance
		})
		resp, err := json.Marshal(issues)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	})
}
