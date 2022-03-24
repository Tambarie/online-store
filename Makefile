run: |
	gofmt -w .
	go run main.go

mock-product:
	mockgen -source=service/store.go -destination=service/product_mock.go -package=service

