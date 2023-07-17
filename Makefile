BIN=bin/raytrace
PNG=test.png
PROF=prof/cpu.prof

loop:
	while true; do \
		go build -o $(BIN) ./src \
		&& time $(BIN) -scene test > tmp.png \
		&& mv tmp.png $(PNG); \
		sleep 2; \
	done

run:
	time go run ./src > $(PNG)

bench:
	go test -o $(BIN).test -cpuprofile $(PROF) -bench . ./src > $(PNG)

prof: bench
	go tool pprof --text $(BIN).test $(PROF)

.PHONY: loop run bench prof
