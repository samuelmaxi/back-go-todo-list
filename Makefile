.PHONY: default run build test docs clean

SHELL:=/bin/bash

default: build-with-run

rule:
	source env.sh && source .env

build-with-run:
	@source .env && go build cmd/main.go && ./main