# The Internet

## Overview
Forked from [Northsec the-internet](https://github.com/nsec/the-internet)

We have to modify some lines to build the application because [lxd](https://github.com/lxc/lxd) has been migrated to [incus](https://github.com/lxc/incus)

Tested with:
- Ubuntu 24.04 - Linux 6.8.0-47-generic #47-Ubuntu
- Go 1.22.2 linux/amd6

## Prerequisites
- [Incus](https://linuxcontainers.org/incus/introduction/)

## Build
- Clone
```
$ git clone https://github.com/linxlunx/the-internet
```
- Init
```
$ go mod init
```
- Build
```
$ go build
```

# Starting the whole thing
Creating an Internet simulation is basically as simple as:
 - the-internet create \<path\>
 - the-internet start

Generate an html/js map of your Internet with:
 - the-internet generate-map \<destination path\>

You can stop the simulation with:
 - the-internet stop

Or create a new one by calling the start command again.

Finally, once you want it all off your disk, you can call:
 - the-internet destroy
