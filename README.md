# go-workshop

ardan lab workshop

# 1. Startup

- Startup an HTTP server.
- Respond to all requests with "Hello world"
- Use Postman or similar to make a request.

# 2. Shutdown

- Catch OS signals
- Gracefully shutdown HTTP server
- Add sleep to demonstrate a long-running request

## Links:

- [The complete guide to Go net/http timeouts](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)
- [So you want to expose Go on the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)

## File Changes:

```
Modified main.go
```

# 3. JSON

- Create a type `Product` with fields for Name, Cost, and Quantity
- Rename the `Echo` handler to `ListProducts`
- Create a slice of Product values with some dummy data.
- Marshal the slice to JSON and write it to the client.
- Use `w.WriteHeader` to explicitly set the response status code.
- Include the Content-Type header so clients understand the response.
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
- See what happens when a nil slice is provided.

## File Changes:

```
Modified main.go
```
