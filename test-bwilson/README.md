# test-bwilson

**Assumptions**: 

- You can ssh to each hosts in `hosts/test-bwilson` as `root`

**Limitations**: 
- `ansible 2.4` is installed on the system running the commands
- all commands are run from the project root directory
- The hosts private Wireguard keys are in an ecrypted yaml file `hosts/group_vars/all/private-keys.vault.yml`. You will need the vault password to run any tasks.

[Ansible vault](https://docs.ansible.com/ansible/2.4/vault.html) isn't the greatest solution, but should work fine for this test project.

## Running

There are 3 tags to run subsets of tasks, or run the entire list by not supplying any tags

Run everything:

```
ansible-playbook playbooks/magnesium.yml --ask-vault-pass
```

See tags available:

```
ansible-playbook playbooks/magnesium.yml --ask-vault-pass --list-tags
```

Run a subset of tags:

```
ansible-playbook playbooks/magnesium.yml --ask-vault-pass --tags wireguard-config,wireguard-test
```

## Project tags

### `wireguard-config`

Configure wireguard interface and peers

### `wireguard-install`

Configure apt repos and install packages

### `wireguard-test`

Installs a simple python3 http server to test connectivity between peers.

## Project layout

### `ansible.cfg`

local client configuration

### `hosts/group_vars/all/`

variables available to all hosts

### `hosts/test-bwilson`

host inventory, connection details and host specific vars.

This isn't a great layout, but should suffice for a test

### `playbooks/magnesium.yml`

The only playbook defined for this test

### `playbooks/roles/wireguard/`

Ansible role to install and configure Wireguard

- `handlers` are tasks intiated by other tasks
- `tasks` the tasks!
- `tempates` resources for tasks to use

