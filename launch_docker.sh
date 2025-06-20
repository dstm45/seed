#! /bin/bash
sudo docker run \
  --name db \
  -e POSTGRES_PASSWORD=12345 \
  -p 5432:5432 \
  -v pgdata:/var/lib/postgresql/data \
  -d postgres