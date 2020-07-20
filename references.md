#### problems and solved
###### 1. Cannot reference internal module when using GO111MODULE=on
In modules, there finally is a name for the subdirectory. If the parent directory go.mod says "module m" then the subdirectory is imported as "m/subdir", no longer "./subdir".
###### 2. Resize centos disk space:
https://support.nagios.com/kb/article/resizing-the-vm-disk-size-for-centos-7-814.html

in last step, using:

> xfs_growfs /dev/root_vg/root
 
instead of

> resize2fs /dev/root_vg/root

### More cmd on centos

- open port on centos
> firewall-cmd --zone=public --permanent --add-port=5000/tcp
firewall-cmd --zone=public --permanent --add-port=4990-4999/udp
firewall-cmd --zone=public --permanent --list-ports

- view cluster dashboard:
> kubectl proxy --address='0.0.0.0' --disable-filter=true

- update container image 
>kubectl rollout restart deployment

- Open kube proxy for test
>kubectl proxy --address='0.0.0.0' --disable-filter=true

### References

Kubernetes infrastructure: https://platform9.com/blog/kubernetes-enterprise-chapter-2-kubernetes-architecture-concepts/

Mongodb on kubernetes: https://switchit-conseil.com/2019/10/16/deploy-a-secured-high-availability-mongodb-replica-set-on-kubernetes/

test on go: https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318

create image from container: https://www.scalyr.com/blog/create-docker-image/

how to test gRPC: https://dzone.com/articles/testing-a-grpc-service-with-table-driven-tests-1

golang standard project: https://github.com/golang-standards/project-layout
