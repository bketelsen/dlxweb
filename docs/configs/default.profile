$> lxc profile show default
config:
  raw.idmap: |
    both 1000 1000
description: Default LXD profile
devices:
  eth0:
    name: eth0
    nictype: bridged
    parent: br0
    type: nic
  keys:
    path: /home/bjk/.ssh
    source: /home/bjk/.ssh
    type: disk
  root:
    path: /
    pool: lxd
    type: disk
name: default
used_by:
- /1.0/instances/dlxweb