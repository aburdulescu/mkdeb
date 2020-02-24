ifeq ($(MKDEB_VERSION),)
	VERSION = dev
else
	VERSION = $(MKDEB_VERSION)-0
endif

ifeq ($(MKDEB_ARCH),)
	ARCH = $(shell dpkg --print-architecture)
else
	ARCH = $(MKDEB_ARCH)
endif

DEB_DIR = mkdeb_$(VERSION)_$(ARCH)
DEB_NAME = $(DEB_DIR).deb

all:
	go build -ldflags "-s -w -X main.version=$(MKDEB_VERSION)"

clean:
	go clean
	rm -rf $(DEB_DIR) $(DEB_NAME)

deb: all
	mkdeb $(DEB_DIR)
	dpkg-deb -b $(DEB_DIR)
