loop:
	while true; do time go run src/*.go > tmp.png && mv tmp.png test.png; sleep 2; done

run:
	go run src/*.go > test.png

.PHONY: loop run
