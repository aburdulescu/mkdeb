all:
	go build -ldflags "-s -w"

clean:
	go clean
	rm -rf mkdeb.out *.deb

deb: all
	./mkdeb
	dpkg-deb -b mkdeb.out
