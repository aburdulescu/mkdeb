$(info $(MKDEB_VERSION))
ifeq ($(MKDEB_VERSION),)
	DEB_DIR = mkdeb-dev
else
	DEB_DIR = mkdeb-$(MKDEB_VERSION)
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
