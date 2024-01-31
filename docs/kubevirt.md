### KubeVirt guidelines

It is expected that you have a kubernetes cluster with KubeVirt installed already.

If you wish to make use of the [CDI](https://kubevirt.io/user-guide/operations/containerized_data_importer/) support
then it should be installed, and you should also have an S3 endpoint to which you can push the image that is built.

Baski will push the image to the S3 endpoint then create a **DataVolume** (and required credentials if they don't exist)
that will reference the image at the S3 endpoint as a source. 

If you wish to manually do this yourself then set `kubevirt.store-in-s3` to false.

Then craft a `baski.yaml` file based on the [example](../baski-example.yaml) supplied and run the commands you require.
