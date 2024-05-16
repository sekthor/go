# Build Go Binary without CGO

This is a test to solve a problem I was having when working with sqlite with golang.
The widely used [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) package complies sqlite from c source code and links the go binary against it.
This is done using `cgo` and will link against the c stdlib.
Therefore it will only run on systems with glibc, musl, or other implementations installed.

The following example will not work:

```Docker
FROM alpine
COPY ./binary ./
CMD ./binary
```

because we still need to install the c stdlib:

```Docker
RUN apk add gcompat
```

## Pure Go Solution

As I am aiming for a statically linked binary without external dependencies we need to stop linking against libc.
The solution is to use pure go implementations of sqlite.
These reimplement sqlite from scratch in golang.

I use [glebarez/go-sqlite](https://github.com/glebarez/go-sqlite).

However initially I was getting the same error.
I was still missing libc.

I realized, that there were some tests, that were using cgo in that package.
It has no impact on the package we use, but are still seen by the compiler.
Therefore we need to prevent CGO from being used at all.

In the compilations step we need to set `CGO_ENABLED` to `0`.

```Docker
RUN CGO_ENABLED=0 go build -o binary .
```