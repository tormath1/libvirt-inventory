## Ansible Libvirt dynamic inventory

Generate an Ansible dynamic inventory (all / active) from libvirt

### Build

```shell
$ git clone https://github.com/tormath1/libvirt-inventory
$ cd libvirt-inventory
$ make
```

### Usage

```shell
$ ansible -m ping -i ./inventory all 
```

### Todo

- [ ] customizable (URI, etc.)
- [ ] add unit tests
- [ ] handle different network interfaces
- [ ] handle differents "tag" using name of the domains ?
