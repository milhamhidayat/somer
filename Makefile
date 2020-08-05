IMAGE_NAME = somer

.PHONY: vendor
vendor: go.mod go.sum
	GO111MODULE=on go get ./...


