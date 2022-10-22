.PHONY serve, build-confs:

build-confs: 
	@cat confs-*.yaml > ./data/confs.yaml

	
serve: build-confs
	@hugo server
