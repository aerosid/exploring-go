#!/bin/bash
set -e
set -x
# create user 'ubuntu'@'localhost' identified by 'hello';
# grant all privileges on *.* to 'ubuntu'@'localhost' with grant option;
# flush privileges;
docker run \
--detach \
--env MYSQL_DATABASE=sample \
--env MYSQL_ROOT_PASSWORD=hello \
--env MYSQL_PASSWORD=hello \
--env MYSQL_USER=ubuntu \
--rm \
--name mysql \
--network host \
--volume /home/ubuntu/vscode/exploring-go/misc/mysql:/var/lib/mysql \
mysql:8.0.30
# sleep 3s
# docker exec -it mysql /bin/bash -c "mysql -uroot -phello"

