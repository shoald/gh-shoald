# gh-shoald

Team tools for the Shoald GitHub org, as a [gh](https://cli.github.com) extension.

## Install

You need `gh` logged in with your Shoald account, and `git`.

```
gh extension install shoald/gh-shoald
```

To update later:

```
gh extension upgrade shoald
```

## Usage

```
gh shoald new repo-name
gh shoald new --template=go repo-name
```

`new` creates a private repo from a template, applies our repo settings, clones it, and fills in the repo name.
Flags go before the name.
Run `gh shoald help` for all commands.

## Development

Build and link your local copy:

```
go build
gh extension remove shoald
gh extension install .
```

After this, `gh shoald` runs the binary in this folder. Run `go build` again after each change.

## Release

Tag a version and push it:

```
git tag v0.1.0
git push origin v0.1.0
```

The release workflow builds binaries for each platform and attaches them to a GitHub release.
