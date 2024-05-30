# TODO sort commands

# list all the local images
docker images
# show image details
docker inspect image_name

# 构建镜像
docker build --tag=<image> path
# search images
docker search image_name

# 从指定仓库拉取镜像
docker pull image_name

# delete image
docker rmi image_name

# list running containers
docker ps
# list all containers
docker ps -as

# delete all containers
docker rm -f 'docker ps -a -q'

docker stop 'docker ps -a -q'  停止所有运行的容器

docker login 登陆

docker pull registry-dev.youle.game/misc/<image>:<tag>

docker tag SOURCE_IMAGE[:TAG] registry-dev.youle.game/misc/IMAGE[:TAG]

docker push registry-dev.youle.game/misc/IMAGE[:TAG] -v /c/Users/webroot:/home/server/webroot

# list all running containers, the same as `docker ps`
# `-a`: List all containers, whatever the status they are.
docker container ls

# delete container
docker container rm [container_id]
# force delete all containers
docker container rm -f $(docker container ls -aq)

docker exec
-it [container_name] /bin/bash	Interact with particular container

docker run
# `-d`: Execute as a daemon process
# `-p`: 3306:3306  端口映射，可在`/var/lib/docker/containers/[ID]/hostconfig.json`下修改
# `-e MYSQL_ROOT_PASSWORD=123456` 		执行命令，初始化mysql密码
# `--name [name]`: claim container name
# -v /home/chy/mysql/config/my.cnf:/etc/mysql/my.cnf 	共享目录挂载


