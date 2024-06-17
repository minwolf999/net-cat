# Net-Cat

Netcat est une implémentation de Netcat en Go, adaptée pour servir de chat avec plusieurs canaux (channels). Chaque canal est limité à 10 utilisateurs maximum.

## Fonctionnalités

- Chat en temps réel avec plusieurs canaux
- Limite de 10 utilisateurs par canal
- Support des protocoles TCP
- Modes client et serveur
- Écoute sur des ports spécifiques
- Envoi et réception de messages

## Prérequis

Go version 1.22.1 ou supérieure

## Installation

Clonez ce dépôt et compilez le projet :
```
git clone https://github.com/minwolf999/net-cat.git
cd net-cat
```

## Utilisation

### Mode Serveur

Utiliser cette commande pour lancer le serveur (par défault le port 8080 est utilisé s'il n'est pas préciser dans la commande):
```
go run . [PORT]
```

### Mode Utilisateur

Pour rejoindre un canal de chat sur le serveur (remplacer IPV4 par l'ipv4 de l'appareil ayant lancer le serveur et PORT par le port utilisé sur l'appareil):
```
nc [IPV4] [PORT]
```