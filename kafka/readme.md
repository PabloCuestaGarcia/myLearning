# Kafka Cluster setup and administration

## AWS

* Setup network security to allow Zookeeper ports (2181, 2888, 3888)
* Setup network security to allow my IP only
* Create EC2 Ubuntu machines with size t2.medium ( 4gb RAM )
* Reserver 3 private IPs for our machines

## Single machine setup

* SSH into our machine
* Install packages on the machine
* Disable RAM Swap
* Add hosts mapping from hostnmae to public ips to /etc/hosts
* Download and configure Zookeeper on the machine
* Launch Zookeeper on the machine to test
* Setup Zookeeper as a service on the machine

