# Install dependencies
---
go mod tidy

# Run the server
---
go run main.go

# Important
---
Change the db url in function getConnection from repository/common.go
