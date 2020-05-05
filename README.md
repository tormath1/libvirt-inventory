## Ansible Libvirt dynamic inventory

![Go](https://github.com/tormath1/libvirt-inventory/workflows/Go/badge.svg)

Generate an Ansible dynamic inventory (all / active) from libvirt

![demo](./doc/inventory.gif)

### Build

Before building, you need at least `libvirt-dev` installed on your machine.

```shell
$ git clone https://github.com/tormath1/libvirt-inventory
$ cd libvirt-inventory
$ make
```

### Install

This step will install the `inventory` into `/etc/ansible/hosts/libvirt-inventory`

```shell
$ make && sudo make install
```

### Usage

If installed:

```shell
$ ansible -m ping all
```

if just built:

```shell
$ ansible -m ping ./inventory all
```

### Todo

- [ ] handle different network interfaces
- [ ] handle differents "tag" using name of the domains ?
