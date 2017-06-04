## Flaq Documentation

### Getting Started

*TODO...*

### Setting Up A New Server

We need to prep our VM before we can run the Ansible playbook.

```bash
$ ssh root@<vm-hostname-or-ip>

# Install Ansible dependencies
$ apt-get update
$ apt-get install ssh python -y

# Create "flaq user
$ sudo adduser flaq # password: "dvs9837B"

# Add "flaq" user to sudo group
$ sudo adduser flaq sudo

# Copy SSH key to "flaq" user account
$ sudo cp -R /root/.ssh /home/flaq/.ssh
$ sudo chown -R flaq:flaq /home/flaq/.ssh

# Quit the SSH session
$ exit
```

You'll need to add the hostname or IP of the new server to `ansible/inventories/prod`.

Now from the root directory of this repository, run the following:
```bash
$ cd ansible
$ ansible-playbook playbook.yml -i inventories/prod -l <vm-hostname-or-ip>
```
