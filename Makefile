.PHONY serve, build-confs, build-contributors, lint:

build-contributors:
	@cp ./contributors.yaml ./data/

build-confs: 
	@cat confs-*.yaml > ./data/confs.yaml

lint:
	@echo "Linting YAML files..."
	@echo "INFO: This will also fail if yamllint is not installed."
	@echo ""
	@yamllint confs-*.yaml config.yaml ./data/confs.yaml -s
	@echo "Lint succeeded!"

	
serve: build-confs build-contributors
	@hugo server
