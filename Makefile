.PHONY: fmt

# Run gofmt but exclude vendor directories
fmt:
	@echo "Running gofmt on package folders $(PKGDIRS)"
	goimports -d $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./Godeps/*")