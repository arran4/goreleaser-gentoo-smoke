# Extra Files Installation in Gentoo ebuild

## Context

In GoReleaser, packages frequently include additional files besides the binary, such as man pages (`.1` files) or configuration files (`.yaml`, `.conf`). These extra files are typically defined in the `archives` section and copied over.

The proposed or implemented Gentoo support uses helper-specific fields corresponding to Gentoo's install functions (e.g., `doins`, `doman`, `dobin`, `doconfd`, `dodir`, `dodoc`, `doenvd`, `doexe`, `doheader`, `doinitd`, `dosbin`, `dosym`, etc.):

```yaml
useflags:
  - flag: systemd
    description: enables systemd installation

files:
  - goreleaser-gentoo-smoke.1
  - config.yaml

doins:
  - src: config.yaml
    dst: /etc/goreleaser-gentoo-smoke/goreleaser-gentoo-smoke.conf

doman:
  - goreleaser-gentoo-smoke.1

doinitd:
  - src: initd-script.sh
    dst: /etc/init.d/goreleaser-gentoo-smoke
    use: [systemd]

dosym:
  - src: /usr/bin/goreleaser-gentoo-smoke
    dst: /usr/bin/ggs
```

This generates Gentoo-native `insinto`/`newins`, `doman`, `doinitd`, and `dosym` commands, properly guarded by `if use <flag>; then ... fi` where `use` arrays are specified. The specified `useflags` are also aggregated into the ebuild's `IUSE` and the package's `metadata.xml`. Meanwhile, `extra_install` remains the escape hatch for advanced installation logic.