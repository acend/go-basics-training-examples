package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var _ IssueGetter = &FileIssueService{}

type FileIssueService struct {
	file string
}

func NewFileIssueService(file string) *FileIssueService {
	return &FileIssueService{
		file: file,
	}
}

func (fis *FileIssueService) GetIssues(_ context.Context) ([]*Issue, error) {
	file, err := os.Open(fis.file)
	if err != nil {
		return nil, fmt.Errorf("failed to open issue file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	issues := []*Issue{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("incorrect number of fields in file")
		}
		importance, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse importance: %w", err)
		}
		issues = append(issues, &Issue{
			Title:       parts[1],
			Description: parts[2],
			Importance:  importance,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read issues: %w", err)
	}

	return issues, nil
}
