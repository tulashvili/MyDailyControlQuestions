APP="My daily control questions"

.PHONY:	run_linter

run_linter:
	golangci-lint run
