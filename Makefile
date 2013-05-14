ROOT := $(CURDIR)

test:
	cd ${ROOT}/4-packages/evens && go test && \
	cd ${ROOT}/4-packages/stack && go test
