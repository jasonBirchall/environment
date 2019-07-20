# environment
Collection of Ansible and bash scripts to setup environments.

## install ansible using ansible repository
To setup your environment you'll need to install ansible and run playbooks.

- Install ansible: 
```
$ sudo apt upgrade
$ sudo apt install software-properties-common
$ sudo apt-add-repository ppa:ansible/ansible
$ sudo apt update
$ sudo apt install ansible
```
- Create a symlink
```
$ sudo ln -s ~/Documents/workarea/environment/ansible/hosts /etc/ansible/hosts
```
