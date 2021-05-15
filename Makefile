TARGET=cyt

build:
	go build -o $(TARGET) .

run:
	go run .

clean:
	rm -f $(TARGET)

.PHONY: run build clean
