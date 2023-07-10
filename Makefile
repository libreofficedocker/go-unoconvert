BUILD_DIR := build

it:
	go mod tidy

run:
	go run cli/unoconvert.go

build:
	$(call build,linux)
	$(call build,darwin)

clean:
	rm -rf $(BUILD_DIR) | true

# define a reusable recipe
define build
	@echo "Building for $(1)..."
	CGO_ENABLED=0 GOOS=$(1) \
		go build -o $(BUILD_DIR)/unoconvert-$(1) cli/unoconvert.go
endef
