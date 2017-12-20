###########################
# Makefile for Controller #
###########################
buildExecutable: 
  $(shell go build -o "out/Controller.exe") \
  $(shell exit 0)
.PHONY: all
all: buildExecutable