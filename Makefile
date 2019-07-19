.PHONY: bench
bench:
	go test -bench=. ./...

.PHONY: bench-off-optimizations
bench-off-optimizations:
	go test -bench=. ./... -gcflags '-N -l'
