#!/bin/sh
apt update
apt install git curl wget -y
apt-add-repository ppa:ansible/ansible
apt update
apt install ansible python-pip -y
curl -fsSL https://get.docker.com | sh
pip install docker-py
