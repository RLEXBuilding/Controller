###########################
# Makefile for Controller #
###########################
all: buildExecutable

buildExecutable:
	go build -o "out/Controller.exe"
.PHONY: all buildExecutable