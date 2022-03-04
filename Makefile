SOURCE=$(find . -name '.go')

mrmorse: $(SOURCE)
	go mod tidy
	go build .

web: mrmorse
	./mrmorse


test:
	go test .
