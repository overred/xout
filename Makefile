help:
	@echo "bench ---------- run all tests and benchmarks"
	@echo "prof-xtarget --- run tests, benchmarks, and save profiles for xtarget"
	@echo "prof-xforamt --- run tests, benchmarks, and save profiles for xformat"
.PHONY: help

bench:
	go test --bench=. --benchtime=1s --benchmem ./...
.PHONY: bench

prof-xout:
	go test --cpuprofile=prof_xout_cpu.prof --memprofile=prof_xout_mem.prof --bench=. --benchtime=1s --benchmem .
.PHONY: prof-xtarget

prof-xtarget:
	go test --cpuprofile=prof_xtarget_cpu.prof --memprofile=prof_xtarget_mem.prof --bench=. --benchtime=1s --benchmem ./xtarget
.PHONY: prof-xtarget

prof-xformat:
	go test --cpuprofile=prof_xformat_cpu.prof --memprofile=prof_xformat_mem.prof --bench=. --benchtime=1s --benchmem ./xformat
.PHONY: prof-xformat

bench:
	go test --cpuprofile=prof_bench_cpu.prof --memprofile=prof_bench_mem.prof --bench=. --benchtime=1s --benchmem ./benchmarks
.PHONY: bench