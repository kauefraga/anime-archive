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
    <img
      alt="GitHub LICENSE"
      src="https://img.shields.io/github/license/kauefraga/anime-archive"
    />
    <a href="https://github.com/kauefraga/anime-archive/commits/main">
      <img
        alt="GitHub last commit"
        src="https://img.shields.io/github/last-commit/kauefraga/anime-archive"
      />
    </a>
  </p>
</div>

> I used to write the animes that i watch right in a plain text file. Therefore, i created Anime Archive to facilitate the process of opening file explorer, searching for the file where i used to write those animes that i watched, write a anime with its url and maybe upload it to Google Drive, just for backup. I decided to use Golang because it's a fast, compiled and a reliable programming language with garbage collector, concurrency and a robust built-in library. See my Python attempt of it, [anime-list](https://github.com/kauefraga/anime-list-python).

### Features

Why would you use Anime Archive instead of a notepad? Here's why:

- [x] A simple interface, you will easily figure out how to use it.
- [x] It's blazingly fast because don't has any server calls. It registers stuff in a file-based database (SQLite).
- [x] It's also lightweight because of the database choice.

And the most important one, if you are a command line enjoyer, then I probably gotcha :).

Need some help? See usage section or run `anime-archive` or `anime-archive --help`

## ⬇️ How to install and use it

1. Install the executable at [releases](https://github.com/kauefraga/anime-archive/releases)
2. As i don't develop a feature to automatically add the application to the PATH, you will need to add it manually.

```bash
# Here is a way to add it to the PATH
# Inside the folder where you have installed Anime Archive...

# Windows (powershell)
$pwd=pwd
$env:PATH="$env:PATH;$pwd" # One line: $pwd=pwd;$env:PATH="$env:PATH;$pwd"

# Linux
export PATH="$PATH:$(pwd)"
```

## 🤹‍♂️ Usage

- To setup Anime Archive, use `setup`.
- To store an anime, use `store "ANIME TITLE" "ANIME URL"`.
- To search for an anime, use `search "ANIME TITLE"`.
- To list all the stored animes, use `list`.
- To list *alternative sites* to watch anime (:brazil:), use `status`.
- To export the database to a human readable format, use `export [--json or --csv]`.

## 🦄 Technical Stuff

If you want to learn more about the project and maybe help me to improve it, see the [ONBOARDING](ONBOARDING.md).

## 📋 Possible improvements

- `list`: add flag functions like tail and head that queries only a specified number of rows
- Create a rest API that reads the database.

## 📝 License

This project is licensed under the GPL-3.0 License - See the [LICENSE](https://github.com/kauefraga/anime-archive/blob/main/LICENSE) for more information.
