#!/usr/bin/env make -f

export GOPATH := $(CURDIR):$(CURDIR)/vendor
export PATH := $(CURDIR)/bin:$(PATH)

.PHONY: devel ./bin/%

default: devel

test: $(wildcard src/**/*.go) $(wildcard vendor/src/**/*.go) ./bin/gb
	gb test

bin/gb-vendor: $(wildcard vendor/src/github.com/constabulary/**/*.go)
	go build -o bin/gb-vendor github.com/constabulary/gb/cmd/gb-vendor

bin/gb: bin/gb-vendor $(wildcard vendor/src/github.com/constabulary/**/*.go)
	go build -o bin/gb github.com/constabulary/gb/cmd/gb

bin/gopherjs: $(wildcard src/**/*.go) $(wildcard vendor/src/**/*.go) ./bin/gb
	gb build github.com/gopherjs/gopherjs

bin/hugo: $(wildcard src/**/*.go) $(wildcard vendor/src/**/*.go) ./bin/gb
	gb build github.com/spf13/hugo

bin/% : $(wildcard src/**/*.go) $(wildcard vendor/src/**/*.go) ./bin/gb
	gb build $( basename $@ )

devel: bin/gen
	gb build gen
	./bin/gen

serve: bin/server
	gb build gen/server
	./bin/server

doc:
	godoc -http :6060

deploy:
	git checkout gh-pages
	./bin/gopherjs build gen/frontend
	cp src/gen/frontend/index.html .
	git add frontend.js
	git add index.html
	git commit -m "Deploying"
	git push -f origin HEAD:gh-pages

