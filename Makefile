ifeq ($(MKDEB_VERSION),)
	DEB_DIR = mkdeb_dev
else
	DEB_DIR = mkdeb_$(MKDEB_VERSION)-0_$(shell dpkg --print-architecture)
endif

DEB_NAME = $(DEB_DIR).deb

all:
	go build -ldflags "-s -w"

clean:
	go clean
	rm -rf $(DEB_DIR) $(DEB_NAME)

deb: all
	./mkdeb $(DEB_DIR)
	dpkg-deb -b $(DEB_DIR)
