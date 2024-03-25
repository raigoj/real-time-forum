clear
docker image build -f Dockerfile -t forum .
sleep 1
docker container run -p 8080:8080 --detach --name docker-forum forum
echo 'http://localhost:8080'