package main

func main() {
	server := NewServer(":3000")
	server.handle("GET", "/", server.AddMiddleware(HandlerRoot, Logging()))
	server.handle("POST", "/home", server.AddMiddleware(HandlerHome, Logging(), CheckAuth()))
	server.handle("POST", "/users", server.AddMiddleware(PostHandlerUser, Logging()))
	server.handle("POST", "/v2/users", server.AddMiddleware(HandlerUser, Logging()))
	server.listen()
}
