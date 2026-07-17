# Missing Support for Extra Files Installation in Gentoo ebuild

## Context

In GoReleaser, packages frequently include additional files besides the binary, such as man pages (`.1` files) or configuration files (`.yaml`, `.conf`). These extra files are typically defined in the `archives` section and copied over.

Different packagers in GoReleaser handle the installation of these extra files in various ways. For instance:
* `brews` provides an `install` block where users can manually specify `man1.install "file.1"` or `etc.install "config.yaml"`.
* `nfpms` uses a `contents` array mapping the source file to an exact destination (`dst`) and setting the `type` (e.g. `man` or `config`).

However, the current implementation in the `feature/gentoo-ebuild` PR (and the `gentoos` configuration) does not provide a mechanism for the generated ebuild to actually *install* extra files that are specified.

While the `Gentoo` struct contains an array of `Files` of type `ExtraFile` which correctly copies the files into the overlay's `files/` directory alongside the ebuild, the generated `src_install()` function in the `ebuildTemplate` only supports installing the compiled binaries and standard documentation (`dodoc README*`). It completely ignores any of the custom extra files copied to the `files/` folder.

## The Gap

The Gentoo `src_install()` function requires the use of helpers like `doman` for man pages, `doins` for general files (with `insinto` to specify the directory), or `doinitd` for init scripts, to properly place extra files into the system tree.

Currently, there is no way for the `.goreleaser.yaml` to instruct the Gentoo pipe to add `doman` or `doins` to the generated ebuild.

## What Needs to be Done

To address this gap, the GoReleaser Gentoo PR needs to be updated. Here is a prompt that can be provided to a developer or agent working on that PR:

---

### Request for the GoReleaser Gentoo PR branch:

The current Gentoo ebuild generation in GoReleaser copies `ExtraFile`s to the `files/` directory in the overlay but lacks a mechanism to actually install them within the `src_install()` function of the generated ebuild. We need a way for the user to configure how these extra files should be installed (e.g., using `doman` for man pages, `doins` for configs, etc.).

Please update the `Gentoo` struct and the ebuild generation logic to support installing extra files. You can look at how `nfpms` handles this with the `contents` array, or how `brews` uses the `install` script section, as inspiration.

**Requirements:**
1. Update `pkg/config/config.go` to add a way to map the extra files to their installation paths or types (e.g., a `contents` field like NFPM, or an `install` script block like Homebrew). If using a structured approach, you might need something like:
   ```yaml
   gentoos:
     - name: my-package
       # ...
       contents:
         - src: my-package.1
           dst: /usr/share/man/man1/my-package.1
           type: man
         - src: config.yaml
           dst: /etc/my-package/config.yaml
           type: config
   ```
2. Update the `ebuildTemplate` in `internal/pipe/gentoo/gentoo.go` to iterate over these configurations in `src_install()` and output the correct Gentoo helpers:
   - Use `doman` for man pages.
   - Use `insinto <dir>` followed by `doins <file>` for arbitrary files.
3. Update the logic in `internal/pipe/gentoo/gentoo.go` to populate the data needed by the template for these extra files.
4. Add tests to ensure that `src_install` properly includes `doman`, `insinto`, and `doins` based on the user's configuration.

Please implement this feature in the `feature/gentoo-ebuild` branch so that the generated ebuild correctly installs man pages and other necessary files.