# Run sample
run:
	go run main.go

# Test
test: build
	ginkgo -r