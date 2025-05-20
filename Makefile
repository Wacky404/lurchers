BINARY=bin/lurchers
SRC=cmd/lurchers/main.go

run: $(BINARY)
	./$(BINARY)

$(BINARY): $(SRC)
	go build -o $(BINARY) $(SRC)
