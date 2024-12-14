
kmds: init
	go build -o _output/kdms cmd/main.go

init:
	mkdir -p _output
