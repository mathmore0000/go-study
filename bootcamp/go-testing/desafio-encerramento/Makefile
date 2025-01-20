# This command will generate a coverage report for the project
.PHONY: coverage
coverage:
	go test -cover -coverprofile=coverage.out ./...

# This command will display the coverage report in the terminal
.PHONY: coverage-total
coverage-total: coverage
	go tool cover -func="coverage.out"