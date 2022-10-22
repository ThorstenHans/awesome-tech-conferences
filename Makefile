.PHONY serve, build-confs, lint:

build-confs: 
	@cat confs-*.yaml > ./data/confs.yaml

lint:
	@echo "Linting YAML files..."
	@echo "INFO: This will also fail if yamllint is not installed."
	@echo ""
	@yamllint confs-*.yaml config.yaml ./data/confs.yaml -s
	@echo "Lint succeeded!"

	
serve: build-confs
	@hugo server
