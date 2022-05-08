package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		debug = false
	)

	// parse command line
	flag.BoolVar(&debug, "debug", debug, "show additional information about rate limiting")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing argument: action")
		os.Exit(1)
	}

	action := flag.Arg(0)
	args := flag.Args()[1:]

	// setup github client
	githubClient := NewClient()
	githubClient.Debug = debug

	if token, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
		githubClient.Token = token
	}

	switch action {
	case "repo":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "missing argument: org and/or name")
			os.Exit(1)
		}

		org := args[0]
		name := args[1]

		repo, err := githubClient.GetRepo(org, name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("name: %s\n", repo.Name)
		fmt.Printf("description: %s\n", repo.Description)
		fmt.Printf("stars: %d\n", repo.Stars)
		fmt.Printf("issues: %d\n", repo.Issues)

	case "user":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "missing argument: username")
			os.Exit(1)
		}

		name := args[0]

		user, err := githubClient.GetUser(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("name: %s\n", user.FullName)
		fmt.Printf("followers: %d\n", user.Followers)

	default:
		fmt.Fprintf(os.Stderr, "unknown action: '%s'\n", action)
		os.Exit(1)
	}
}
