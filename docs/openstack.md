### OpenStack guidelines

It is expected that you have a network and sufficient security groups in place to run this.<br>
It will not create the network or security groups for you.

For example:

```
openstack network create image-builder
openstack subnet create image-builder --network image-builder --dhcp --dns-nameserver 1.1.1.1 --subnet-range 10.10.10.0/24 --allocation-pool start=10.10.10.10,end=10.10.10.200
openstack router create image-builder --external-gateway public1
openstack router add subnet image-builder image-builder

OS_SG=$(openstack security group list -c ID -c Name -f json | jq '.[]|select(.Name == "default") | .ID')
openstack security group rule create "${OS_SG}" --ingress --ethertype IPv4 --protocol TCP --dst-port 22 --remote-ip 0.0.0.0/0 --description "Allows SSH access"
openstack security group rule create "${OS_SG}" --egress --ethertype IPv4 --protocol TCP --dst-port -1 --remote-ip 0.0.0.0/0 --description "Allows TCP Egress"
openstack security group rule create "${OS_SG}" --egress --ethertype IPv4 --protocol UDP --dst-port -1 --remote-ip 0.0.0.0/0 --description "Allows UDP Egress"
```

Then craft a `baski.yaml` file based on the [example](../baski-example.yaml) supplied and run the commands you require.

## Image Visibility

When scanning an image, if the `infra.openstack.image-visibility` field is not set, then Baski will make it public if
the scan succeeds. This means that all projects in OpenStack can access this image.

If the scan fails, it'll be made private - even if it was previously public, because insecure images shouldn't be
available.

By supplying the `infra.openstack.image-visibility` field when running a scan, you can override this behaviour.

The recommendation is that when building and scanning an image for the first time, the field is left blank. This way a
new, insecure image won't be made available.

If you're using the `scan.multiple.image-search`, then it's recommended to set the visibility (presuming they are always
the same - for example, **public**) as Baski will try to override the visibility otherwise.

⚠️ **_This could break things such as autoscaling in Kubernetes on a cluster that is still using the now insecure image._**

In production, it's prudent to get that cluster onto a newer, secure build of the image asap but in reality I know this
can take time hence the ability to override.
