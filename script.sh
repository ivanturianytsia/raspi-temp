function container_run {
  docker container run \
    --rm \
    --env PORT=3000 \
    --name temp \
    -p 3000:3000 \
    -d \
    ivanturianytsia/raspi-temp
}

container_run
