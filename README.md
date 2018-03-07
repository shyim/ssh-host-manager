# SSH Host Manager

[![Greenkeeper badge](https://badges.greenkeeper.io/shyim/ssh-host-manager.svg)](https://greenkeeper.io/)

Simple CLI to manage your ssh shortcuts

## Install

```bash
npm install -g ssh-host-manager
```

## Requirements

* Unix system
* NodeJs
* OpenSSH with min version 7.3p1

## Usage

```bash
// Add a new entry with name test
shm add test root@example.com

// Add a new entry with name test with custom port
shm add test root@example.com:443

// Add a new entry with name test with custom ssh key
shm add test root@example.com /home/shyim/.ssh/server.key

// Connect to the new entry
ssh test
```

![List](https://ipfs.io/ipfs/QmWDyiBECcKC2A8EADkK1N7bCKNmoY9ovi6d6cN6z2VbLa)

## Uninstall

* Remove the file
``$HOME/.ssh/manager_hosts``

* Remove the include in ``$HOME/.ssh/config``
