#!/bin/sh
apt update
apt install git curl wget -y
apt-add-repository ppa:ansible/ansible
apt update
apt install ansible -y
curl -fsSL https://get.docker.com | sh
git clone https://github.com/foxutech/go-app-deployment-on-aws.git
cd go-app-deployment-on-aws
ansible-playbook deployment.yml -vvvv -i localhost >> /home/ubuntu/ansible.txt
