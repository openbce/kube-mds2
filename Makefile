
kmds: init
	go build -o _output/kube-mds2 cmd/main.go

init:
	mkdir -p _output
