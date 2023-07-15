BIN=bin/raytrace
PNG=test.png

loop:
	while true; do go build -o $(BIN) ./src && time $(BIN) -scene random > tmp.png && mv tmp.png $(PNG); sleep 2; done

run:
	time go run ./src > $(PNG)

bench:
	cd src; go test -cpuprofile ../prof/cpu.prof -bench . > $(PNG)

prof:
	go tool pprof --text prof/src.test prof/cpu.prof

.PHONY: loop run bench prof
