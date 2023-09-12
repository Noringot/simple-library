# Simple library

## Overview
A simple application for authors and the books they wrote.
P.S. This repository contains .env and config.yml files, which in normal development it is not recommended to push to a public repository. Here they are for easier deployment of the project

## To start the project you need:
- Golang version 1.20.4+
- Docker
- Protoc

## Installation
### Environment 
In terminal
```bash
git clone git@github.com:Noringot/simple-library.git
```
When project is download
```bash
cd simple-library
make -f makefile gen-docker
```
After docker container is launch
```bash
make -f makefile gen-db
make -f makefile gen-proto
```
### To launch a service
```bash
make -f makefile start-server
```
