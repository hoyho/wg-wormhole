.PHONY: all binary

binary:
	@ echo; echo "### $@:"
	go build -o wormhole ./cmd

install-node: binary
	@ echo; echo "### $@:"
	cp wormhole /usr/local/bin/
	cp deploy/node.service /usr/lib/systemd/system/wormhole-node.service
	sudo systemctl daemon-reload
	sudo systemctl restart wormhole-node
	sudo systemctl enable wormhole-node.service

install-registry: binary
	@ echo; echo "### $@:"
	cp wormhole /usr/local/bin/
	cp deploy/registry.service /usr/lib/systemd/system/wormhole-registry.service
	sudo systemctl daemon-reload
	sudo systemctl restart wormhole-registry
	sudo systemctl enable wormhole-registry.service

clean:
	@ echo; echo "### $@:"
	rm wormhole

code-gen:
	@ echo; echo "### $@:"
	bash proto/getCode.sh