
Vagrant.configure("2") do |config|

  config.vm.box = "ubuntu/xenial64"

  config.vm.network "private_network", ip: "192.168.43.43"

  config.vm.synced_folder ".", "/home/vagrant/go/src/github.com/nicklaw5/flaq", :nfs => { :mount_options => ["dmode=777","fmode=777"] }

  config.vm.provider "virtualbox" do |vb|
    vb.gui = false
    vb.memory = 4096
    vb.cpus = 4
  end

  config.hostmanager.enabled = true
  config.hostmanager.manage_host = true
  config.hostmanager.ignore_private_ip = false
  config.hostmanager.include_offline = true

  config.vm.define "flaq-dev" do |node|
    node.vm.hostname = "dev.flaq.live"
  end

  config.vm.provision "bootstrap", type: "shell" do |s|
    s.inline = "apt-get install python -y"
  end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "ansible/playbook.yml"
    ansible.inventory_path = "ansible/inventories/dev"
    ansible.limit = "all"
  end

  config.vm.provision "file", source: "~/.gitconfig", destination: ".gitconfig"
end
