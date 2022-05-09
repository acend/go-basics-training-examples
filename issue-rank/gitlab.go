package main

import (
	"context"

	"github.com/xanzy/go-gitlab"
)

type GitlabIssueService struct {
	client *gitlab.Client
	repo   string
}

func NewGitlabIssueService(url, repo, token string) (*GitlabIssueService, error) {
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		return nil, err
	}
	return &GitlabIssueService{
		client: client,
		repo:   repo,
	}, nil
}

func (gis *GitlabIssueService) GetIssues(ctx context.Context) ([]*Issue, error) {
	gitlabIssues, _, err := gis.client.Issues.ListProjectIssues(gis.repo, &gitlab.ListProjectIssuesOptions{})
	if err != nil {
		return nil, err
	}
	issues := []*Issue{}
	for _, issue := range gitlabIssues {
		issues = append(issues, &Issue{
			Title:       issue.Title,
			Description: issue.Description,
			Importance:  issue.Upvotes,
		})
	}
	return issues, nil
}
