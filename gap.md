# Missing Support for Extra Files Installation in Gentoo ebuild

## Context

In GoReleaser, packages frequently include additional files besides the binary, such as man pages (`.1` files) or configuration files (`.yaml`, `.conf`). These extra files are typically defined in the `archives` section and copied over.

Different packagers in GoReleaser handle the installation of these extra files in various ways. For instance:
* `brews` provides an `install` block where users can manually specify `man1.install "file.1"` or `etc.install "config.yaml"`.
* `nfpms` uses a `contents` array mapping the source file to an exact destination (`dst`) and setting the `type` (e.g. `config` or omitting it for regular files).

However, the current implementation in the `feature/gentoo-ebuild` PR (and the `gentoos` configuration) does not provide a mechanism for the generated ebuild to actually *install* extra files that are specified.

While the `Gentoo` struct contains an array of `Files` of type `ExtraFile` which correctly copies the files into the overlay's `files/` directory alongside the ebuild, the generated `src_install()` function in the `ebuildTemplate` only supports installing the compiled binaries and standard documentation (`dodoc README*`). It completely ignores any of the custom extra files copied to the `files/` folder.

## The Gap

The proposed or implemented Gentoo support uses helper-specific fields:

```yaml
files:
  - goreleaser-gentoo-smoke.1
  - config.yaml

doins:
  - src: config.yaml
    dst: /etc/goreleaser-gentoo-smoke/goreleaser-gentoo-smoke.conf

doman:
  - goreleaser-gentoo-smoke.1
```

This generates Gentoo-native `insinto`/`newins` and `doman` commands, while `extra_install` remains the escape hatch for advanced installation logic.