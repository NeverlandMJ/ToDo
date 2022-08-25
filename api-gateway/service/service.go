package service

type Provider struct {
	UserServiceProvider
	TodoServiceProvider
}

type UserServiceProvider interface {
	
}

type TodoServiceProvider interface {
	
}

func NewProvider(userServiceURL, todoServiceURL string) Provider {
	return Provider{
		UserServiceProvider: NewGRPCClientUser(userServiceURL),
		TodoServiceProvider: NewGRPCClientTodo(todoServiceURL),
	}
}
