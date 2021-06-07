
.PHONY: all clean gen

all: gen
	@$(MAKE) -C cmd/quack

gen:
	@go generate ./...

clean:
	@$(MAKE) -C cmd/quack clean
