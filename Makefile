.PHONY: build
build:
	@wire && \
	go mod tidy && \
	go build -o ./build/app .

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: run
run:
	@./build/app