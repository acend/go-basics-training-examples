package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Issue struct {
	Title       string
	Description string
	Importance  int
}

type IssueGetter interface {
	GetIssues(ctx context.Context) ([]*Issue, error)
}

func main() {
	var (
		source     = "gitlab"
		file       = "issues.txt"
		gitlabURL  = "https://gitlab.com"
		gitlabRepo = "dvob/issues"
		enableJSON = false
	)
	flag.StringVar(&source, "source", source, "Issue source. Valid values are gitlab, file and dummy.")
	flag.StringVar(&gitlabURL, "gitlab-url", gitlabURL, "URL of the Gitlab instance. Set GITLAB_TOKEN if it is a private repository.")
	flag.StringVar(&gitlabRepo, "gitlab-repo", gitlabRepo, "Gitlab Repositories to read the issues from.")
	flag.StringVar(&file, "file", file, "Read issues from file")
	flag.BoolVar(&enableJSON, "json", enableJSON, "Return output as JSON")
	flag.Parse()

	var (
		err         error
		issueGetter IssueGetter
	)
	switch source {
	case "file":
		issueGetter = NewFileIssueService(file)
	case "gitlab":
		issueGetter, err = NewGitlabIssueService(gitlabURL, gitlabRepo, os.Getenv("GITLAB_TOKEN"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "dummy":
		issueGetter = &DummyIssueService{
			Issues: []*Issue{
				{
					Importance:  0,
					Title:       "Dummy Title",
					Description: "Dummy Description",
				},
			},
		}
	default:
		fmt.Printf("unknown source: '%s'\n", source)
		os.Exit(1)
	}

	var handler http.Handler
	if enableJSON {
		handler = jsonIssueHandler(issueGetter)
	} else {
		handler = textIssueHandler(issueGetter)
	}

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
