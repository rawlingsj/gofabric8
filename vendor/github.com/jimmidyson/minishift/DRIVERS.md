# Driver plugin installation

Minikube uses Docker Machine to manage the Kubernetes VM so it benefits from the
driver plugin architecture that Docker Machine uses to provide a consistent way to
manage various VM providers. Minikube embeds VirtualBox and VMware Fusion drivers
so there are no additional steps to use them. However, other drivers require an
extra binary to be present in the host PATH.

The following drivers currently require driver plugin binaries to be present in
the host PATH:

* [KVM](#kvm-driver)
* [xhyve](#xhyve-driver)

#### KVM driver

Minikube is currently tested against `docker-machine-driver-kvm` 0.7.0.

From https://github.com/dhiltgen/docker-machine-kvm#quick-start-instructions:

```
$ sudo curl -L https://github.com/dhiltgen/docker-machine-kvm/releases/download/v0.7.0/docker-machine-driver-kvm -o /usr/local/bin/docker-machine-driver-kvm
$ sudo chmod +x /usr/local/bin/docker-machine-driver-kvm
```

On Debian/Ubuntu
```
# Install libvirt and qemu-kvm on your system, e.g.
$ sudo apt install libvirt-bin qemu-kvm

# Add yourself to the libvirtd group (may vary by linux distro) so you don't need to sudo
$ sudo usermod -a -G libvirtd $(whoami)

# Update your current session for the group change to take effect
$ newgrp libvirtd
```

On Fedora
```
# Install libvirt and qemu-kvm on your system, e.g.
$ sudo dnf install libvirt qemu-kvm

# Add yourself to the libvirt group so you don't need to sudo
$ sudo usermod -a -G libvirt $(whoami)

# Update your current session for the group change to take effect
$ newgrp libvirt
```

#### xhyve driver

From https://github.com/zchee/docker-machine-driver-xhyve#install:

```
$ brew install docker-machine-driver-xhyve

# docker-machine-driver-xhyve need root owner and uid
$ sudo chown root:wheel $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve
$ sudo chmod u+s $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve
```
