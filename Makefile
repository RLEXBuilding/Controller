buildExecutable: 
  $(shell go build -o "out/Controller.exe") \
  $(shell go build -o "out/Controller.o") \
.PHONY: all
all: buildExecutable