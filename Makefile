run: |
	gofmt -w .
	go run main.go

mock-service:
	mockgen -source=domain/service/wallet.go -destination=domain/service/mock_wallet_engine_db.go -package=service

