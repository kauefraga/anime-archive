# Onboarding Documentation

> I created this to help you get familiar with Anime Archive. Feel free to open issues and pull requests!

## ⚒ Build

To build Anime Archive you must have: Golang (>=1.20) and GCC.

```bash
# Build for a specific platform
# See: https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
GOARCH={architecture} GOOS={platform} go build -o {outputfile} ./src/main.go

# Default build
make build
```

## 🔨 Scripts

Anime Archive is using `make` and currently has three scripts:

- `build`: as it says, builds the project.
- `release`: builds the project with some optimization flags (`-ldflags=-w`).
- `clean`: runs `go clean` and removes the generated binaries.

Perhaps you want to see the [`Makefile`](Makefile).

## ⚗ Tests

I actually don't know how to write tests for it. 🤡

## 🗃 Database

> I choose [SQLite](https://www.sqlite.org/index.html) because of its power and availability. In fact, i want it to work offline and to be portable.

You need to have GCC or some C compiler to work with SQLite.

## 🛠 Technologies

Here is a list of what packages and technologies i am using to build Anime Archive:

- [Golang](https://go.dev) - A powerful and robust compiled programming language.
- [SQLite](https://www.sqlite.org/index.html) - It's a tiny, fast, self-contained and reliable SQL database engine.
- [GORM](https://gorm.io) - The fantastic ORM library for Golang.
  - [SQLite driver](https://pkg.go.dev/gorm.io/driver/sqlite)
- [Go Pretty](https://github.com/jedib0t/go-pretty) - "Utilities to prettify console output of tables, lists, progress-bars, text, etc. with a heavy emphasis on customization."
- [XDG](https://pkg.go.dev/github.com/adrg/xdg) - Go implementation of the XDG Base Directory Specification and XDG user directories.
- [Cobra](https://cobra.dev) - A framework for modern CLI apps in Go.
