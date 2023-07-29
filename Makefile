BIN=bin/raytrace
PNG=test.png
TEST=bin/raytrace.test
PROF=prof/cpu.prof

SCENE=test

loop:
	while true; do \
		go build -o $(BIN) ./internal \
		&& time $(BIN) -samples 16 -threads 2 -scene $(SCENE) > tmp.png \
		&& mv tmp.png $(PNG); \
		sleep 2; \
	done

run:
	time go run ./internal > $(PNG)

scene_test:
	time go run ./internal -samples 32 -scene test > $(PNG)

scene_random:
	time go run ./internal -samples 64 -scene random > $(PNG)

scene_light:
	time go run ./internal -scene light > $(PNG)

scene_cornell:
	time go run ./internal -samples 100 -aspect 1.0 -scene cornell > $(PNG)

scene_smoke:
	time go run ./internal -samples 200 -aspect 1.0 -scene smoke > $(PNG)

scene_final:
	time go run ./internal -samples 1000 -aspect 1.0 -scene final > $(PNG)


bench:
	go test -c -o $(TEST) ./internal && \
	$(TEST) -test.cpuprofile $(PROF) -test.bench . > $(PNG) 

prof: bench
	go tool pprof --text $(TEST) $(PROF)

.PHONY: loop run bench prof
