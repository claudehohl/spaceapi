all:
	go build spaceapi.go
	./spaceapi

raspi:
	GOARCH=arm GOARM=7 go build spaceapi.go

.PHONY: all
