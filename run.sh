#!/bin/bash
set -e
set -x
docker run \
--detach \
--rm \
--name go \
--network host \
--user ubuntu:ubuntu  \
--volume /etc/timezone:/etc/timezone:ro \
--volume /etc/localtime:/etc/localtime:ro \
--volume /home/ubuntu/.git-credentials:/home/ubuntu/.git-credentials:ro \
--volume /home/ubuntu/.gitconfig:/home/ubuntu/.gitconfig:ro \
--volume /home/ubuntu/vscode/exploring-go:/home/ubuntu/vscode/exploring-go \
--workdir /home/ubuntu/vscode/exploring-go \
go:latest tail -f /dev/null
sleep 3s
docker exec -it go bash