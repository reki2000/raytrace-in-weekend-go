BIN=bin/raytrace
PNG=test.png
TEST=bin/raytrace.test
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
	go test -c -o $(TEST) ./src && \
	$(TEST) -test.cpuprofile $(PROF) -test.bench . > $(PNG) 

prof: bench
	go tool pprof --text $(TEST) $(PROF)

.PHONY: loop run bench prof
