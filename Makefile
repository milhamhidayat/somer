IMAGE_NAME = somer

# Go Modules
.PHONY: vendor
vendor: go.mod go.sum
	GO111MODULE=on go get ./...

.PHONY: tidy
tidy:
	go mod tidy

