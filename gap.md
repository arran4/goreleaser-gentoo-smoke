# Extra Files Installation in Gentoo ebuild

## Context

In GoReleaser, packages frequently include additional files besides the binary, such as man pages (`.1` files) or configuration files (`.yaml`, `.conf`). These extra files are typically defined in the `archives` section and copied over.

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