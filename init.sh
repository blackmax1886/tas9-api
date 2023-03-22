# !/bin/bash

# wait for Mysql server starting
until mysqladmin ping -h tas9-db -P 3306 --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting...!"
exec ./main
