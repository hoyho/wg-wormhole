.PHONY: all binary

binary:
	@ echo; echo "### $@:"
	go build -o wormhole ./cmd

clean:
	@ echo; echo "### $@:"
	rm wormhole

code-gen:
	@ echo; echo "### $@:"
	bash proto/getCode.sh