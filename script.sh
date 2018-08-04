function container_run {
  docker pull ivanturianytsia/raspi-temp:latest

  docker container run \
    --env PORT=3000 \
    --env DATA_PATH=/data \
    -v temp:/data \
    --name temp \
    -p 3000:3000 \
    -d \
    --rm \
    -u root:root \
    ivanturianytsia/raspi-temp
}

container_run
