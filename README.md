<div align="center">
  <h1><code>Anime Archive</code></h1>

  <p>
    <strong>🦋 A command line interface to create, find and list all my viewed animes. 🦋</strong>
  </p>

  <p>
    <img
      alt="GitHub top language"
      src="https://img.shields.io/github/languages/top/kauefraga/anime-archive.svg"
    />
    <img
      alt="Repository size"
      src="https://img.shields.io/github/repo-size/kauefraga/anime-archive.svg"
    />
    <a href="https://github.com/kauefraga/anime-archive/commits/main">
      <img
        alt="GitHub last commit"
        src="https://img.shields.io/github/last-commit/kauefraga/anime-archive.svg"
      />
    </a>
    <img
      alt="GitHub LICENSE"
      src="https://img.shields.io/github/license/kauefraga/anime-archive.svg"
    />
  </p>
</div>

> I used to write the animes that i watch right in a plain text file. Therefore, i created Anime Archive to facilitate the process of opening file explorer, searching for the file where i used to write those animes that i watched, write a anime with its url and maybe upload it to Google Drive, just for backup. I decided to use Golang because it's a fast, compiled and a reliable programming language with garbage collector, concurrency and a robust built-in library. See my Python attempt of it, [anime-list](https://github.com/kauefraga/anime-list-python).

### Features

- A fancy UI that is easy to use.
- Everything looks better with some colors, and i am addicted to it.
- It uses SQLite, the whole database is a file.
- If you forgot how to use it, just ask for help: `anime-archive --help`.
- Intuitive commands: `setup`, `store`, `search`, `list`, `export` and `status`.

## ⬇️ How to install and use it

1. Install the executable at [releases](https://github.com/kauefraga/anime-archive/releases)
2. As i don't develop a feature to automatically add the application to the PATH, you will need to add it by yourself.

```bash
# Here is a way to add it to the PATH
# Inside the folder where you have installed Anime Archive...

# Windows (powershell)
$pwd=pwd
$env:PATH="$env:PATH;$pwd" # One line: $pwd=pwd;$env:PATH="$env:PATH;$pwd"

# Linux
export PATH="$PATH:$(pwd)"
```

## 🦄 Technical Stuff

If you want to learn more about the project and maybe help me to improve it, see the [ONBOARDING](ONBOARDING.md).

## 📋 Possible improvements

- Create a rest API that reads the database.

## 📝 License

This project is licensed under the MIT License - See the [LICENSE](https://github.com/kauefraga/anime-archive/blob/main/LICENSE) for more information.
