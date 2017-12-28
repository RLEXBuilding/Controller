###########################
# Makefile for Controller #
###########################
buildExecutable:
	go build -o "out/Controller.exe"
.PHONY: all buildExecutable
all: buildExecutable

# TODO:
# 		- Multios Build