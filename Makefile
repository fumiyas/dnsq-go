GO=		go
GOX=		gox

DIRS=		build/bin
PACKAGE_ROOT=	github.com/fumiyas/dnsq-go
CMD_PKGS=	./cmd/dnsq ./cmd/dnsqr

CROSS_TARGETS=	linux/amd64 darwin/amd64 windows/amd64

default: build

dirs:
	mkdir -p $(DIRS)
	
build: dirs
	$(GO) build -o ./cmd/bin/dnsq ./cmd/dnsq
	$(GO) build -o ./cmd/bin/dnsqr ./cmd/dnsqr

cross: dirs
	$(GOX) \
	  -osarch='$(CROSS_TARGETS)' \
	  -output='./build/bin/{{.Dir}}.{{.OS}}.{{.Arch}}' \
	  $(CMD_PKGS) \
	;

clean:
	$(RM) -rf $(DIRS)
