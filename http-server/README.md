# HTTP Server Exercises


## General

1. Run a server and implement a handler which returns `pong` on the endpoint `/ping`.
2. Create a middleware to log every request. Log the path, method and duration of the request.


## User API

* Create a JSON REST-like API where you can create, list, get and delete users.

The API should provide the following endpoints:

* `POST /user`: create a user
* `GET /user`: returns all users
* `GET /user/<ID>`: returns a single user by ID
* `DELETE /user/<ID>`: delete user by ID

The user should look at least like this:
```
{
    "name": "john",
    "full_name": "John Doe",
    "followers": 13
}
```

You can decide how you would like to store the users. One option would be to store them only in memory for example in a map `map[int]User`. Another option would be to store serialize the users into a file (e.g. with JSON).

* Create the appropriate client implementation to create users

For this you could write a CLI tool which takes the username as argument and the full name and the followers as option:

```
./usercli alice
./usercli bob -full "Bob Miller" -followers 34
```

* Add HTTP basic authentication to the whole user API

You can do this best if you wrap the API handlers in a appropriate middleware. For the authentication you can use a static username password pair like `admin` and `secret`. You can also try to make them configurable using flags or environment variables.
