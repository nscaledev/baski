# Changelog

## [ 2024/03/xx - v1.2.0 ]

### BREAKING CHANGES

### Remove
* Baski server has been removed - it was extra work that was sporadically used. It's up to the user to provide a way to
  interact with the results. If you wish to keep server, then look at a previous release and fork that code by all means.

## [ 2024/03/11 - v1.1.1 ]

### Changed/Added

* Setting images to private by default and will only set to public once a scan has passed successfully - this can still
  be overridden at the config level.
  
## [ 2024/02/29 - v1.1.0 ]

### BREAKING CHANGES

* changed `cloud` prefix to `infra` in flags and config.
* changed `build.nvidia` prefix to `build.gpu` in flags and config.
* changed `build.nvidia.enable-nvidia-support` prefix to `build.gpu.enable-gpu-support` in flags and config.

### Changed/Added

* Added KubeVirt as a build option.
* Supports AMD GPUs

## [ 2024/02/15 - v1.0.0 ]

### Changed/Added

First release with:

* Functioning support for OpenStack build, scan and signing.
* Baski Server 
