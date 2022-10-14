docker network create kafka
docker run --name server01 -p 8080:8080 --network kafka -dit ubuntu:22.04 bash
docker run --name server02 -p 8090:8080 --network kafka -dit ubuntu:22.04 bash
docker run --name server03 -p 8091:8080 --network kafka -dit ubuntu:22.04 bash

# Update packages
apt update && apt install -y wget nano openjdk-8-jdk

wget https://dlcdn.apache.org/kafka/3.2.0/kafka_2.12-3.2.0.tgz
tar -xf kafka_2.12-3.2.0.tgz && rm kafka_2.12-3.2.0.tgz
mkdir /kafka
nano /kafka-3.2.0/config/zookeper.properties
nano /kafka-3.2.0/config/server.properties