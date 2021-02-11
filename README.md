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