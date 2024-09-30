FILE := "main.go"
TARGET := "main"

default:
	go build $(FILE)

run:
	./main

clean:
	rm $(TARGET)
	
