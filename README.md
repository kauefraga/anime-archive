<div align="center">
  <h1><code>Anime Archive</code></h1>

  <p>
    <strong>🦋 A command line interface to create, find and list all my viewed animes. 🦋</strong>
  </p>

  <p>
    <img
      alt="GitHub top language"
      src="https://img.shields.io/github/languages/top/kauefraga/anime-archive"
    />
    <img
      alt="GitHub all releases"
      src="https://img.shields.io/github/downloads/kauefraga/anime-archive/total"
    />
    <a href="https://github.com/kauefraga/anime-archive/commits/main">
      <img
        alt="GitHub last commit"
        src="https://img.shields.io/github/last-commit/kauefraga/anime-archive"
      />
    </a>
    <img
      alt="GitHub LICENSE"
      src="https://img.shields.io/github/license/kauefraga/anime-archive"
    />
  </p>
</div>

### Features

Why would you use Anime Archive instead of a notepad? Here's why:

- [x] A simple interface, you will easily figure out how to use it.
- [x] It's blazingly fast because it doesn't have any server calls. It registers stuff in a file-based database (SQLite).
- [x] It's also lightweight because of the database choice.

And the most important one, if you are a command line enjoyer, then I probably gotcha :).

Need some help? See the usage section or run `anime-archive` or `anime-archive --help`

## ⬇️ How to install and use it

1. Install the executable at [releases](https://github.com/kauefraga/anime-archive/releases)
2. Add the Anime Archive binary to the PATH

```bash
# Inside the folder where you have installed Anime Archive...

# Windows (powershell)
$pwd=pwd
$env:PATH="$env:PATH;$pwd" # One line: $pwd=pwd;$env:PATH="$env:PATH;$pwd"

# Linux
export PATH="$PATH:$(pwd)"
```

## 🤹‍♂️ Usage

- To setup Anime Archive, use `setup [--useRemoteDatabase]`.
  - `--useRemoteDatabase (-u)` install [my personal anime list](https://github.com/kauefraga/anime-archive/blob/main/animes.db).
- To register an anime, use `register "ANIME TITLE" "ANIME URL"`.
  - `--description (-d) string` assign an description about the anime.
- To search for an anime, use `search "ANIME TITLE"`.
- To list all the stored animes, use `list [--tail uint or --head uint]`.
  - `--head (-H) uint` query some of the newest records.
  - `--tail (-T) uint` query some of the oldest records.
- To list *alternative sites* to watch anime (:brazil:), use `status`.
- To export the database to a human readable format, use `export [--json or --csv]`.
  - `--json (-j)` export the database in JSON format.
  - `--csv (-c)` export the database in CSV format.

## 🦄 Technical Stuff

If you want to learn more about the project and maybe help me to improve it, see the [ONBOARDING](ONBOARDING.md).

## 📜 History

> I used to write the animes that i watch right in a plain text file. Therefore, i created Anime Archive to facilitate the process of opening file explorer, searching for the file where i used to write those animes that i watched, write a anime with its url and maybe upload it to Google Drive, just for backup. I decided to use Golang because it's a fast, compiled and a reliable programming language with garbage collector, concurrency and a robust built-in library, also because i am learning it. See my Python attempt of it, [anime-list](https://github.com/kauefraga/anime-list-python).

## 📝 License

This project is licensed under the GPL-3.0 License - See the [LICENSE](https://github.com/kauefraga/anime-archive/blob/main/LICENSE) for more information.
