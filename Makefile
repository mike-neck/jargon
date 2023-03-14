PATHS:=$(GOPATH) $(HOME)/go
XGEN:=$(strip $(firstword $(foreach p,$(PATHS),$(wildcard $(p)/bin/xgen))))

.PHONY: show
show:
	@echo "$(PATHS)"
	@echo "$(HOME)"

ifeq ($(XGEN),)
.PHONY: install
install:
	@go get github.com/xuri/xgen
	@go install github.com/xuri/xgen/cmd/...
else
.PHONY: xgen
xgen:
	@"$(XGEN)" -i ./data -o ./xml -p main  -l Go
	@mv "$$(find xml -type f -name '*.go' | head -n 1)" ./pom.go
endif
