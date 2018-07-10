code-gen:
	./hack/update-codegen.sh

prereq:
	go get -u \
		github.com/kubernetes/gengo/examples/deepcopy-gen

clean:
	$(RM) $(PROG_BIN)

.PHONY: all clean prereq code-gen