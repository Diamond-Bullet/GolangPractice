# 查看本地所有镜像
docker images

docker inspect ... 查看镜像细节

# 构建镜像
docker build --tag=<image> path

search ... 搜索镜像

# 从指定仓库拉取镜像
docker pull ...

rmi ... 删除镜像

ps 查看运行中的容器

ps -as 查看所有容器

docker rm -f `docker ps -a -q` 删除所有容器

docker stop `docker ps -a -q`  停止所有运行的容器

docker login 登陆

docker pull registry-dev.youle.game/misc/<image>:<tag>

docker tag SOURCE_IMAGE[:TAG] registry-dev.youle.game/misc/IMAGE[:TAG]

docker push registry-dev.youle.game/misc/IMAGE[:TAG] -v /c/Users/webroot:/home/server/webroot