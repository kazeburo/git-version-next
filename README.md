# git-version-next

This package is build based on https://github.com/b4b4r07/git-bump

git-version-next increments version (git tag) numbers simply. Unlike git-bump, it only displays the new version to STDOUT.

## Installation

Download from GitHub Releases and drop it in your $PATH.

## Usage

```
Usage:
  git-version-next [OPTIONS]

Application Options:
      --patch  update patch version
      --minor  update minor version
      --major  update major version

Help Options:
  -h, --help   Show this help message
```

run `git version-next`

```
% git version-next
Use the arrow keys to navigate: ↓ ↑ → ← 
? Current tag is 0.2.10. Next is: 
    patch update (0.2.11)
  ▸ minor update (0.3.0)
    major update (1.0.0)
```

choose next version and enter

```
% git version-next
✔ minor update (0.3.0)
0.3.0

```

## Example usage

Implement bump function.

```
function bump {
  local NEXT=$(git version-next)
  if [ -z $NEXT ]; then
    exit 1
  fi
  git tag v$NEXT
  git push origin v$NEXT
}
```
