#!/bin/bash

# Espera até que o Kafka esteja disponível
echo "Esperando o Kafka ficar disponível..."
while ! nc -z kafka 29092; do
  sleep 1
done
echo "Kafka está disponível."

# Criando tópicos Kafka
echo "Criando tópicos Kafka..."
kafka-topics --bootstrap-server kafka:29092 --create --if-not-exists --topic go-to-rails --replication-factor 1 --partitions 1
kafka-topics --bootstrap-server kafka:29092 --create --if-not-exists --topic rails-to-go --replication-factor 1 --partitions 1
echo "Tópicos Kafka criados com sucesso:"
kafka-topics --bootstrap-server kafka:29092 --list
