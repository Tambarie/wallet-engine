run: |
	gofmt -w .
	go run main.go

mock-service-user:
	mockgen -source=service/user.go -destination=service/user_mock.go -package=service

