#!/bin/bash

# Enable ssh password authentication
echo "[TASK 1] Enable ssh password authentication"
sed -i 's/^PasswordAuthentication .*/PasswordAuthentication yes/' /etc/ssh/sshd_config
echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config
systemctl reload sshd

# Set Root password
echo "[TASK 2] Set root password"
echo -e "admin\nadmin" | passwd root >/dev/null 2>&1


echo "K0s Kubernetes distrubution installation"
curl -sSLf https://get.k0s.sh | sudo sh
k0s install controller --single
k0s start