# Simple Regex Matcher

This is a simple regex matcher implementation in Go, created for educational purposes. It supports basic regex patterns like `^`, `$`, `?`, `*`, and `+`.

## Usage

To use the regex matcher, simply run the `main` function and provide the regex pattern and target string separated by a pipe (`|`) character. For example:

```
go run main.go 'a*b|aaaab'
```

This will match the regex pattern `a*b` against the target string `aaaab` and print `true` if it matches, or `false` otherwise.

## Benchmarks

The project includes benchmarks to compare the performance of the custom regex matcher against Go's built-in `regexp` package. Here are the benchmark results:

```
BenchmarkMatchComplex-8 33666723 35.43 ns/op

BenchmarkRegexPackageComplex-8 7817226 152.6 ns/op

BenchmarkMatchLargeInput-8 37611 31343 ns/op

BenchmarkRegexPackageLargeInput-8 7846782 155.2 ns/op
```

As you can see, the custom regex matcher performs better for complex patterns but lags behind the built-in `regexp` package for large inputs.
