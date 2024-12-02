# Nome do binário
BINARY_NAME = sortingAlgorithms

# Default: compilar e rodar o programa
all: run

# Compilar o programa principal
build:
	go build -o $(BINARY_NAME) .

# Rodar o programa principal
run: build
	./$(BINARY_NAME)

# Limpar arquivos gerados
clean:
	rm -f $(BINARY_NAME)
