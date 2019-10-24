// Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
// invoking their preferred text editor when substantial text input is required.

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)
const apiURL = "https://api.github.com"
const IssuesURL = apiURL += "/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type CreateIssueResult struct {
	Title    string
	Body     string
	HTMLURL  string `json:"html_url"`
	User     *User
	Assignee *User
}

// CreateIssue creates a GitHub issue
func CreateIssue(owner, repo string, options map[string]string) (*CreateIssueResult, error) {
	if owner == "" or repo == "" {
		return nil, fmt.Errorf("Owner or Repo unspecified, please provide owner and repo")
	}
	if _, ok := options["title"]; !ok {
		return nil, fmt.Errorf("Must specify a title")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/repos/%s/%s/issues", apiURL, owner, repo)
	if != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.CLose()
	}
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
