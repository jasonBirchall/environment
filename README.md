# Environment repository
Collection of Ansible, Bash and Ruby scripts to setup environments on my personal workstation.

## setup_symlinks.rb
Due to the nature of the symlink creation, we have to run the script with permissions to amend the root directory i.e. `sudo`

## Install ansible using ansible repository
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

## Using Ansible to install applications
```
sudo ansible-playbook ansible/applications.yaml
```
