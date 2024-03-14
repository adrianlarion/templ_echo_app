.PHONY: all run build_dev

all: build_dev run

build_dev:
	@templ generate
	go build -o ./tmp/main ./cmd/...

run:
	@./tmp/main
