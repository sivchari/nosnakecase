# nosnakecase
nosnakecase is a linter that detects snake case of variable naming and function name.

## Instruction

```sh
go install github.com/sivchari/nosnakecase/cmd/nosnakecase@latest
```

## Usage

```go
package sandbox

func a_() {
}

func b(a_a int) {
}

func c() (c_c int) {
    c_c = 1
	return c_c
}

func d() {
	var d_d int
	_ = d_d // It's never detected, because `_` is meaningful in Go and `d_d` is already detected.
}
```

```console
go vet -vettool=(which nosnakecase) ./...

# command-line-arguments
./main.go:3:6: a_ is used under score. You should use mixedCap or MixedCap.
./main.go:6:8: a_a is used under score. You should use mixedCap or MixedCap.
./main.go:10:2: c_c is used under score. You should use mixedCap or MixedCap.
./main.go:11:9: c_c is used under score. You should use mixedCap or MixedCap.
./main.go:9:11: c_c is used under score. You should use mixedCap or MixedCap.
./main.go:15:6: d_d is used under score. You should use mixedCap or MixedCap.
```

## CI

### CircleCI

```yaml
- run:
    name: install nosnakecase
    command: go install github.com/sivchari/nosnakecase/cmd/nosnakecase@latest

- run:
    name: run nosnakecase
    command: go vet -vettool=`which nosnakecase` ./...
```

### GitHub Actions

```yaml
- name: install nosnakecase
  run: go install github.com/sivchari/nosnakecase/cmd/nosnakecase@latest

- name: run nosnakecase
  run: go vet -vettool=`which nosnakecase` ./...
```
