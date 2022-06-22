help:
	@echo "bench ---------- run all tests and benchmarks"
	@echo "prof-xtarget --- run tests, benchmarks, and save profiles for xtarget"
.PHONY: help

bench:
	go test --bench=. --benchtime=1s --benchmem ./...
.PHONY: bench

prof-xtarget:
	go test --cpuprofile=prof_xtarget_cpu.prof --memprofile=prof_xtarget_mem.prof --bench=. --benchtime=1s --benchmem ./xtarget
.PHONY: prof-xtarget