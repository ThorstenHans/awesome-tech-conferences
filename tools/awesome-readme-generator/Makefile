.PHONY: build copy-to-repo-root

build:
	@echo "Building..."
	@go build -ldflags="-X 'github.com/ThorstenHans/awesome-tech-conferences/tools/awesome-readme-generator/cmd/awesome-readme-generator.version=0.0.1'" -o bin/awesome-readme cmd/awesome-readme-generator/main.go

copy-to-repo-root:
	@echo "Copying binary to repository root"
	@cp ./bin/awesome-readme ../../
