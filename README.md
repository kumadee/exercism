[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/kumadee/exercism)

# exercism
Solutions for pratice problems from https://exercism.io/my/tracks/go

To run tests for all packages, just run the below command:
```bash
go test -v -bench . -benchmem ./...
```

To run the CPU and memory profiling, run the tests with the `cpuprofile` and `memprofile` flags.
```bash
# Assuming we are already in the directory with go.mod
for package in $(find go -maxdepth 1 -mindepth 1 -type d)
do
    go test -v -bench . -benchmem -cpuprofile ${dir}/cpu.out -memprofile ${dir}/mem.out ./${dir}
done
```

To view the profile data in browser, run the below command.
```bash
go tool pprof -http=:8080 cpu.out
go tool pprof -http=:8080 mem.out
```