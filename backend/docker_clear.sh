docker stop $(docker ps -l -a -q)
docker rm docker-forum
docker rmi forum