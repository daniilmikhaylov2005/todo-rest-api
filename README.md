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

## /auth/

>SignUp 

username and password can't be empty


**takes**

```yaml
"username": "someUsername",
"name": "someName"
"password": "somepassword"
```
**return**

```yaml
"id: "6",
"status": "User created"
```

>SignIn

**takes**

```yaml
"username": "someUsername",
"password": "somepassword"
```

**return**

```yaml
"token": "someJwtTokee like a dsakljdaokjd1323.13dasdas.13123das"
```

## /api/

to access for this route you need to use jwt token
role for Create, Update, Delete operations is admin
basic role after registration is user
User can only get all todos and get todo by id

>GetAllTodos

**takes**

Only username from jwt token

**return**

```yaml
[{"id": 8,
 "userId": 5,
"title": "go to shop",
"done": false}, ...]
```

>GetTodoById

**takes**

Id param from url and username from jwt token

**return**

```yaml
"id": 8,
"userId": 5,
"title": "go to shop",
"done": false
```

>CreateTodo

**takes**

role from jwt token and id param from url

```yaml
"title": "sometitle",
"done": false
```

**return**

```yaml
"id": 15,
"userId": 5,
"title": "somebody did",
"done": false
```

>UpdateTodo

When you change "done" u can don't add "title", example:

```yaml
"done": true
```


**takes**

role from jwt token and id param from url

```yaml
"title": "updatedTitle",
"done": true
```

**return**

Updated todo

```yaml
"id":16,
"userId":5,
"title":"updatedTitle",
"done": true
```

>DeleteTodo

**takes**

role from jwt token and id param from url

**return**

```yaml
"status": "todo with id {id} deleted"
```
