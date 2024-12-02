BINARY_NAME = sortingAlgorithms

all: run

build:
	go build -o $(BINARY_NAME) ./cmd

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
