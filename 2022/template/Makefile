MAKEFLAGS += --no-print-directory

default: main.go input.go
	go run main.go input.go

test: main.go input.go main_test.go
	go test main_test.go main.go input.go

watch: main.go input.go
	while true; \
		do inotifywait main.go 2>/dev/null && \
		clear && make; \
	done

watch-test: main.go input.go main_test.go
	while true; \
		do inotifywait main_test.go 2>/dev/null && \
		clear && make test; \
	done
