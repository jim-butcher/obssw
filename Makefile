build: obssw obscamtoggle

obssw:
	go build ./cmd/obssw/

obssw-clean:
	rm obssw
	go clean ./cmd/obssw/

obscamtoggle:
	go build ./cmd/obscamtoggle/

obscamtoggle-clean:
	rm obscamtoggle
	go clean ./cmd/obscamtoggle/

clean : obssw-clean obscamtoggle-clean
	go clean .
