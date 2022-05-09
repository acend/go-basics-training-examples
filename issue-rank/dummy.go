package main

import "context"

var _ IssueGetter = &DummyIssueService{}

// DummyIssueService serves Ideas from memory
type DummyIssueService struct {
	Issues []*Issue
}

func (dis *DummyIssueService) GetIssues(_ context.Context) ([]*Issue, error) {
	return dis.Issues, nil
}
