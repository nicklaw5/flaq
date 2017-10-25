# flaq *(working title)*

## Requirements

- Virtual Box
- Vagrant
- Vagrant Host Manager
- Ansible (v2.3.1 or higher)

## Development Environment

### Server Setup

Follow the below commands to configure and start the streaming server.

```bash
# Clone the repository
$ git clone git@github.com:nicklaw5/flaq.git
$ cd flaq

# Install Ansible Galaxy dependencies
$ ansible-galaxy install -r ansible/requirements.yml --roles-path=ansible/roles

# (Ubuntu users only) Add firewall exception for our VM. Otherwise NFS access to the VM is unavailable.
$ sudo ufw allow from 192.168.43.0/24

# Fire up the VM (grab a coffee - this takes roughly 5-10 mins)
$ vagrant up
```

### OBS Setup

To start streaming to the server from OBS, follow these steps:

1. With OBS running, click on ***Settings***
2. Select ***Stream*** from the left-hand menu
3. Change ***Stream Type*** to ***Custom Streaming Server***
4. In the ***URL*** field, copy and past the following: `rtmp://dev.flaq.live/app`
5. In the ***Stream key*** field, copy and past the following: `DevStreamKey1`
6. Click ***Apply*** and ***OK*** to save the changes
7. Select ***Start Streaming*** to start streaming to the server

Once you've started streaming, you can watch the stream by opening the `index.html` file in the root of this project in Firefox (other browsers are not supported yet)

## Third-party Documentation

Information and documentation for other third-party packages that are used in this repository are listed below:

- [nginx-rtmp wiki](https://github.com/arut/nginx-rtmp-module/wiki)
