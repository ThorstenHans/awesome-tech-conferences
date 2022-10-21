.PHONY serve, build-confs:

build-confs: 
	@cat confs-*.yaml > ./data/confs.yaml

	
serve:
	@hugo server
