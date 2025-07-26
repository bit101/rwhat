build:
	@go build -o rwhat main.go
	@./rwhat -urlmDb --progress --stats --delete /foo / bar
