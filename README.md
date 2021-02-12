[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/kumadee/exercism)

# exercism
Solutions for pratice problems from https://exercism.io/my/tracks/go

To run tests for all modules, just run the below command:
```bash
for file in $(find go -name "go.mod" -type f); do
    pushd $(dirname ${file});
    go test -v --bench . --benchmem;
    popd;
done
```

To run the CPU and memory profiling, run the tests with the `cpuprofile` and `memprofile` flags.
```bash
# Assuming we are already in the directory with go.mod
go test -v --bench . --benchmem -cpuprofile cpu.out -memprofile mem.out
```

To view the profile data in browser, run the below command.
```bash
go tool pprof -http=:8080 cpu.out
go tool pprof -http=:8080 mem.out
```