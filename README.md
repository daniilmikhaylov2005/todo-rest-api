# Install dependencies
go mod tidy

# Run the server
go run main.go

# Important

Change the db url in function getConnection from repository/common.go

# API

## Endpoints:handlerName:Method
> /auth/signin:SignIn:POST
> /auth/signup:SignUp:POST
> /api/todo:GetAllTodos:GET
> /api/todo:CreateTodo:POST
> /api/todo/{id}:GetTodoById:GET
> /api/todo/{id}:UpdateTodo:PUT
> /api/todo/{id}:DeleteTodo:DELETE

## Work with data in api (JSON)
>SignUp 
**takes*
```yaml
		"username": "someUsername",
		"name": "someName"
		"password": "somepassword"
```
**return**
```yaml
		"token": "someJwtTokee like a dsakljdaokjd1323.13dasdas.13123das"
```
