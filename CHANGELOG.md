# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.1] - 2025-03-21

### :gear: Changed
- Just adjusting the codeowners and pining softprops/action-gh-release to a sha by @drew-viles in [#34](https://github.com/nscaledev/baski/pull/34)


## [1.3.0] - 2025-03-21

### :gear: Changed
- Changing git cliff to bump mor than patch by @drew-viles
- Moving ownership to nscale and bumping modules by @drew-viles in [#33](https://github.com/nscaledev/baski/pull/33)
- Bump github.com/go-jose/go-jose/v4 from 4.0.4 to 4.0.5 in the go_modules group across 1 directory by @dependabot[bot] in [#30](https://github.com/nscaledev/baski/pull/30)

### :rocket: Added
- Adding rocky support, updating module, removing dockerfile builds & minor cleanup/refactoring by @drew-viles in [#32](https://github.com/nscaledev/baski/pull/32)


## New Contributors
* @nscale-actions[bot] made their first contribution
## [1.2.5] - 2025-01-13

### :bug: Fixed
- Fixing trivy panics when ignoreList is len 0 by @drew-viles

### :gear: Changed
- Updating how git cliff works in the repo with examples in the README by @drew-viles
- Updating go modules by @drew-viles
- Updating how options are created and handled in the scanner and signer provisioners by @drew-viles
- Changing how concurrent image scanning works to improve error checking by @drew-viles

### :rocket: Added
- Removing all flags references from application calls by @drew-viles


## [1.2.4] - 2024-12-12

### :gear: Changed
- Updating go modules by @drew-viles in [#22](https://github.com/nscaledev/baski/pull/22)


## [1.2.3] - 2024-10-21

### :gear: Changed
- Module and golang update by @drew-viles in [#18](https://github.com/nscaledev/baski/pull/18)

### :rocket: Added
- Adding metadata prefix option to signing by @drew-viles in [#15](https://github.com/nscaledev/baski/pull/15)


## [1.2.2] - 2024-09-09

### :bug: Fixed
- Correcting when the k8s metadata should be applied by @drew-viles in [#13](https://github.com/nscaledev/baski/pull/13)

### :gear: Changed
- Updated pipeline and switched to storing images in DockerHub by @drew-viles in [#14](https://github.com/nscaledev/baski/pull/14)
- Metadata prefix + deprecations by @drew-viles in [#11](https://github.com/nscaledev/baski/pull/11)
- Updating OS and trivy support by @drew-viles
- Module updates by @drew-viles in [#10](https://github.com/nscaledev/baski/pull/10)
- Bump github.com/hashicorp/go-retryablehttp from 0.7.5 to 0.7.7 in the go_modules group across 1 directory by @dependabot[bot] in [#9](https://github.com/nscaledev/baski/pull/9)

### :rocket: Added
- Add option to use existing local checkout of image-builder by @yankcrime in [#12](https://github.com/nscaledev/baski/pull/12)


## New Contributors
* @yankcrime made their first contribution in [#12](https://github.com/nscaledev/baski/pull/12)
* @dependabot[bot] made their first contribution in [#9](https://github.com/nscaledev/baski/pull/9)
## [1.2.1] - 2024-04-26

### :gear: Changed
- Go module updates and added contained version override support by @drew-viles in [#8](https://github.com/nscaledev/baski/pull/8)


## [1.2.0] - 2024-03-20

### :bug: Fixed
- Build section of tag pipeline still referenced old path to main.go by @drew-viles

### :gear: Changed
- Updated CHANGELOG by @drew-viles

### :rocket: Added
- Added more linting and testing - updated helm charts push endpoint by @drew-viles in [#6](https://github.com/nscaledev/baski/pull/6)

### :wastebasket: Removed
- Removed baski server by @drew-viles in [#7](https://github.com/nscaledev/baski/pull/7)


## [1.1.1] - 2024-03-11

### :gear: Changed
- Updated go modules & readme by @drew-viles in [#4](https://github.com/nscaledev/baski/pull/4)

### :rocket: Added
- Added information on Image Visibility to the OpenStack docs by @drew-viles in [#5](https://github.com/nscaledev/baski/pull/5)
- Image-visibility changes on a scan pass or fail. by @drew-viles in [#3](https://github.com/nscaledev/baski/pull/3)


## [1.1.0] - 2024-02-29

### :rocket: Added
- Major overhaul to enable multiple builders to be supported in the all phases by @drew-viles in [#2](https://github.com/nscaledev/baski/pull/2)


## [1.0.0] - 2024-02-15

### :rocket: Added
- Initial Commit and Release by @drew-viles


## New Contributors
* @drew-viles made their first contribution
[1.3.1]: https://github.com/nscaledev/baski/compare/v1.3.0..v1.3.1
[1.3.0]: https://github.com/nscaledev/baski/compare/v1.2.5..v1.3.0
[1.2.5]: https://github.com/nscaledev/baski/compare/v1.2.4..v1.2.5
[1.2.4]: https://github.com/nscaledev/baski/compare/v1.2.3..v1.2.4
[1.2.3]: https://github.com/nscaledev/baski/compare/v1.2.2..v1.2.3
[1.2.2]: https://github.com/nscaledev/baski/compare/v1.2.1..v1.2.2
[1.2.1]: https://github.com/nscaledev/baski/compare/v1.2.0..v1.2.1
[1.2.0]: https://github.com/nscaledev/baski/compare/v1.1.1..v1.2.0
[1.1.1]: https://github.com/nscaledev/baski/compare/v1.1.0..v1.1.1
[1.1.0]: https://github.com/nscaledev/baski/compare/v1.0.0..v1.1.0

<!-- generated by git-cliff -->
