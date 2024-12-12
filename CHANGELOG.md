# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
## [unreleased]

### Changed


- Updating go modules
- CHANGELOG updated using git cliff


## [1.2.3] - 2024-10-21

### Changed


- Module and golang update

* adding a comment around how baski fails as a todo
* updated go modules and version


### Fixed


- Adding metadata prefix option to signing


## [1.2.2] - 2024-09-09

### Changed


- Bump github.com/hashicorp/go-retryablehttp from 0.7.5 to 0.7.7 in the go_modules group across 1 directory (#9)
- Module updates

* security updates for modules
- Updating OS and trivy support
- Updated pipeline and switched to storing images in DockerHub

* updated pipeline and switched to storing images in DockerHub
* adding a CODEOWNERS
- Add option to use existing local checkout of image-builder

Handy when you're building locally and don't want Baski fetching a copy
of the image-builder repo each and every time.

Co-authored-by: Drew Hudson-Viles <drew@hudson-viles.uk>


### Fixed


- Metadata prefix + deprecations

* deprecated almost all of the flags in favor of the config file
* optional metadata prefix for images
* updated modules
* cleanup of error printing in tests
* updating go in the pipeline
- Correcting when the k8s metadata should be applied


## [1.2.1] - 2024-04-26

### Changed


- Go module updates and added contained version override support


## [1.2.0] - 2024-03-20

### Changed


- Added more linting and testing - updated helm charts push endpoint

* added more linting and testing - updated helm charts push endpoint
- Removed baski server

* Removed baski server
* Update golang version in pipeline
* Added a default toe the flag for scan-bucket
- Updated CHANGELOG


### Fixed


- Build section of tag pipeline still referenced old path to main.go


## [1.1.1] - 2024-03-11

### Added


- Image-visibility changes on a scan pass or fail.


### Changed


- Updated go modules & readme
- Added information on Image Visibility to the OpenStack docs


## [1.1.0] - 2024-02-29

### Added


- Major overhaul to enable multiple builders to be supported in the all phases


## [1.0.0] - 2024-02-15

### Added


- Initial Commit and Release


