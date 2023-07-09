loop:
	while true; do time go run ./src -scene test > tmp.png && mv tmp.png test.png; sleep 2; done

run:
	go run ./src > test.png

bench:
	cd src; go test -cpuprofile ../prof/cpu.prof -bench .

prof:
	go tool pprof --text prof/src.test prof/cpu.prof

.PHONY: loop run bench prof
