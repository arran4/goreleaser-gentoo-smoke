# goreleaser-gentoo-smoke

Disposable smoke-test application for testing Gentoo ebuild generation in a
GoReleaser pull request.

This repository is intentionally minimal. The Go program only prints its own
name and version. The actual test is the release workflow plus the
`.goreleaser.yaml` configuration.

## Purpose

This repository verifies that a custom GoReleaser build can:

1. build a tiny Go binary,

2. create a GitHub release from a tag,

3. generate a Gentoo binary ebuild,

4. publish the generated ebuild into a separate Gentoo overlay repository,

5. keep only the newest configured ebuild versions.

The overlay repository is:

<https://github.com/arran4/goreleaser-gentoo-smoke-overlay>

## GoReleaser PR under test

This smoke test targets:

<https://github.com/arran4/goreleaser/pull/2>

The release workflow builds GoReleaser from the PR branch rather than using an
official GoReleaser release.

## Repository roles

This repository is the source application and release driver.

The generated ebuilds should be committed to:

<https://github.com/arran4/goreleaser-gentoo-smoke-overlay>

Generated ebuilds should not be committed manually to this repository.

## Required GitHub secret

Create a fine-grained GitHub personal access token with contents read/write
access to:

- `arran4/goreleaser-gentoo-smoke-overlay`

Add it to this repository as an Actions secret named:

```text
GENTOO_OVERLAY_TOKEN
```

The default `GITHUB_TOKEN` is used for releases in this repository.
`GENTOO_OVERLAY_TOKEN` is used for writing generated ebuilds to the overlay
repository.

## Running the smoke test

Commit and push this repository, then push release tags:

```sh
git tag v0.1.0
git push origin v0.1.0
git tag v0.2.0
git push origin v0.2.0
git tag v0.3.0
git push origin v0.3.0
```

Each tag should trigger the release workflow.

## Expected result

The app repo should get GitHub releases for each pushed tag.

The overlay repo should receive generated ebuilds under:

```text
app-misc/goreleaser-gentoo-smoke-bin/
```

After releasing `v0.1.0`, `v0.2.0`, and `v0.3.0`, the overlay repo should
contain only:

```text
app-misc/goreleaser-gentoo-smoke-bin/goreleaser-gentoo-smoke-bin-0.2.0.ebuild
app-misc/goreleaser-gentoo-smoke-bin/goreleaser-gentoo-smoke-bin-0.3.0.ebuild
```

The old `0.1.0` ebuild should have been removed because `.goreleaser.yaml`
sets:

```yaml
keep_versions: 2
```

## Cleanup

This repository is disposable. Tags, releases, generated artifacts, and overlay
ebuilds may be deleted after testing.
