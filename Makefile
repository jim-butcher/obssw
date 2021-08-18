BUILDPATH ?= build/
BUILDFLAGS ?= -ldflags "-w" -v -o $(BUILDPATH)

.PHONY: build clean obssw-clean obscamtoggle-clean

build: $(BUILDPATH)/obssw $(BUILDPATH)/obscamtoggle

$(BUILDPATH)/obssw:
	go build $(BUILDFLAGS) ./cmd/obssw/

obssw-clean:
	go clean ./cmd/obssw/

$(BUILDPATH)/obscamtoggle:
	go build $(BUILDFLAGS) ./cmd/obscamtoggle/

obscamtoggle-clean:
	go clean ./cmd/obscamtoggle/

clean : obssw-clean obscamtoggle-clean
	rm -rf $(BUILDPATH)
	go clean .
