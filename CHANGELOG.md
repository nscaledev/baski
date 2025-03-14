# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.5] - 2025-01-13

### Changed
- Updating how git cliff works in the repo with examples in the README
- Updating go modules
- Updating how options are created and handled in the scanner and signer provisioners
- Changing how concurrent image scanning works to improve error checking
- Removing all flags references from application calls

### Fixed
- Fixing trivy panics when ignoreList is len 0

## [1.2.4] - 2024-12-12

### Changed
- Updating go modules by @drew-viles in [#22](https://github.com/drewbernetes/baski/pull/22)

## [1.2.3] - 2024-10-21

### Added
- Adding metadata prefix option to signing by @drew-viles in [#15](https://github.com/drewbernetes/baski/pull/15)

### Changed
- Module and golang update by @drew-viles in [#18](https://github.com/drewbernetes/baski/pull/18)

## [1.2.2] - 2024-09-09

### Added
- Add option to use existing local checkout of image-builder by @yankcrime in [#12](https://github.com/drewbernetes/baski/pull/12)

### Changed
- Updated pipeline and switched to storing images in DockerHub by @drew-viles in [#14](https://github.com/drewbernetes/baski/pull/14)
- Metadata prefix + deprecations by @drew-viles in [#11](https://github.com/drewbernetes/baski/pull/11)
- Updating OS and trivy support by @drew-viles
- Module updates by @drew-viles in [#10](https://github.com/drewbernetes/baski/pull/10)
- Bump github.com/hashicorp/go-retryablehttp from 0.7.5 to 0.7.7 in the go_modules group across 1 directory by @dependabot[bot] in [#9](https://github.com/drewbernetes/baski/pull/9)

### Fixed
- Correcting when the k8s metadata should be applied by @drew-viles in [#13](https://github.com/drewbernetes/baski/pull/13)

## New Contributors
* @yankcrime made their first contribution in [#12](https://github.com/drewbernetes/baski/pull/12)
* @dependabot[bot] made their first contribution in [#9](https://github.com/drewbernetes/baski/pull/9)
## [1.2.1] - 2024-04-26

### Changed
- Go module updates and added contained version override support by @drew-viles in [#8](https://github.com/drewbernetes/baski/pull/8)

## [1.2.0] - 2024-03-20

### Added
- Added more linting and testing - updated helm charts push endpoint by @drew-viles in [#6](https://github.com/drewbernetes/baski/pull/6)

### Changed
- Updated CHANGELOG by @drew-viles

### Fixed
- Build section of tag pipeline still referenced old path to main.go by @drew-viles

### Removed
- Removed baski server by @drew-viles in [#7](https://github.com/drewbernetes/baski/pull/7)

## [1.1.1] - 2024-03-11

### Added
- Added information on Image Visibility to the OpenStack docs by @drew-viles in [#5](https://github.com/drewbernetes/baski/pull/5)

### Changed
- Updated go modules & readme by @drew-viles in [#4](https://github.com/drewbernetes/baski/pull/4)
- Image-visibility changes on a scan pass or fail. by @drew-viles in [#3](https://github.com/drewbernetes/baski/pull/3)

## [1.1.0] - 2024-02-29

### Changed
- Major overhaul to enable multiple builders to be supported in the all phases by @drew-viles in [#2](https://github.com/drewbernetes/baski/pull/2)

## [1.0.0] - 2024-02-15

### Changed
- Initial Commit and Release by @drew-viles

## New Contributors
* @drew-viles made their first contribution
[1.2.5]: https://github.com/drewbernetes/baski/compare/v1.2.4..v1.2.5
[1.2.4]: https://github.com/drewbernetes/baski/compare/v1.2.3..v1.2.4
[1.2.3]: https://github.com/drewbernetes/baski/compare/v1.2.2..v1.2.3
[1.2.2]: https://github.com/drewbernetes/baski/compare/v1.2.1..v1.2.2
[1.2.1]: https://github.com/drewbernetes/baski/compare/v1.2.0..v1.2.1
[1.2.0]: https://github.com/drewbernetes/baski/compare/v1.1.1..v1.2.0
[1.1.1]: https://github.com/drewbernetes/baski/compare/v1.1.0..v1.1.1
[1.1.0]: https://github.com/drewbernetes/baski/compare/v1.0.0..v1.1.0

<!-- generated by git-cliff -->
