MAKEFLAGS += --no-print-directory

default: main.go input.go
	go run main.go input.go

watch: main.go input.go
	while true; \
		do inotifywait main.go 2>/dev/null && \
		clear && make; \
	done
