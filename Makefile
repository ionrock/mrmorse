SOURCE=$(find . -name '.go')

mrmorse: $(SOURCE)
	go build .

web: mrmorse
	./mrmorse
