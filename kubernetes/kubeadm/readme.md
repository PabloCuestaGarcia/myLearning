# Install Kubernetes Cluster using kubeadm

Follow this documentation to set up a Kubernetes cluster on **Ubuntu 22.04 LTS**.
This documentation guides you in setting up a cluster with one master node and one worker node.

## Assumptions

| Role | FQDN | IP | OS | RAM | CPU 
|-|-|-|-|-|-|
| Master | kmaster.example.com | 172.16.16.100 | Ubuntu 22.04 | 2G | 2
| Worker | kworker.example.com | 172.16.16.101 | Ubuntu 22.04 | 1G | 1

## On both Kmaster and Kworker

Perform all the commands as root user unless otherwise specified

#### Disable firewall

```shell
ufw disable
```

#### Disable swap

```shell
swapoff -a
sed -i '/swap/d' /etc/fstab
```

#### Update sysctl settings for Kubernetes networking

```shell
cat >>/etc/sysctl.d/kubernetes.conf<<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

# Check system
sysctl --system
```