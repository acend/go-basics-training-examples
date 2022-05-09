package main

import (
	"fmt"
	"os"
	"strconv"
	api "user-api"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	args := os.Args[1:]

	if len(args) < 1 {
		return fmt.Errorf("missing argument action")
	}

	action := args[0]
	args = args[1:]

	client := api.NewClient()

	switch action {
	case "create":
		if len(args) < 1 {
			return fmt.Errorf("missing argument: username")
		}
		user := &api.User{
			Name: args[0],
		}
		if len(args) > 1 {
			user.FullName = args[1]
		}
		if len(args) > 2 {
			followers, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}
			user.Followers = followers
		}
		fmt.Println(user)
		return client.Create(user)
	case "get":
		// list
		if len(args) == 0 {
			users, err := client.List()
			if err != nil {
				return err
			}
			for _, user := range users {
				fmt.Printf("%s full_name=%s followers=%d\n", user.Name, user.FullName, user.Followers)
			}
			return nil
		}

		// single user
		userName := args[0]
		user, err := client.Get(userName)
		if err != nil {
			return err
		}
		fmt.Printf("%s full_name=%s followers=%d\n", user.Name, user.FullName, user.Followers)
		return nil
	case "delete":
		if len(args) < 1 {
			return fmt.Errorf("missing argument: username")
		}

		userName := args[0]

		return client.Delete(userName)
	default:
		return fmt.Errorf("unknown action '%s'", action)
	}
}
