build: obssw

obssw:
	go build ./cmd/obssw/

obssw-clean:
	rm obssw
	go clean ./cmd/obssw/

clean : obssw-clean
	go clean .
