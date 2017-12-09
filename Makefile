buildExecutable: 
  go build -o "out/Controller.exe" "./src";
  go build -o "out/Controller.o" "./src";
  done; exit;
all: buildExecutable