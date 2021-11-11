GOFILES := $(wildcard *.go *.yml *.mod *.sum)

TERRA:
	cp $(GOFILES) terra
	go build -o api terra/*.go

KAVA:
	cp $(GOFILES) kava
	go build -o api kava/*.go
  
COSMOS:
	cp $(GOFILES) cosmos
	go build -o api cosmos/*.go
  
SOLANA:
	cp $(GOFILES) solana
	go build -o api solana/*.go

POLYGON:
	cp $(GOFILES) polygon
	go build -o api polygon/*.go
  
VELAS:
	cp $(GOFILES) velas
	go build -o api velas/*.go

BITSONG:
	cp $(GOFILES) bitsong
	go build -o api bitsong/*.go
  
CERTIK:
	cp $(GOFILES) certik
	go build -o api certik/*.go
  
IRIS:
	cp $(GOFILES) iris
	go build -o api iris/*.go
  
POLKADOT:
	cp $(GOFILES) polkadot
	go build -o api polkadot/*.go

GAME:
	cp $(GOFILES) polkadot
	go build -o api polkadot/*.go

REGEN:
	cp $(GOFILES) regen
	go build -o api regen/*.go
