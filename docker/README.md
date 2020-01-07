
## Kafka
docker-compose -f clsuter_kafka.yml up -d

## TO DO LIST
 - Cluster Redis/Redis Sentinel via Docker-compose
 - Cluster ELK via Docker-compose
 - Déploiement Jenkins et interfaçage avec GitHub
 - Déploiement Varnish / Apache / PHP / MySQL via Docker-Compose
 - Playbook déploiement PostGreSQL 10
 

Chemins à créer :
/app/php/conf.d /app/httpd/conf.d /app/redis/conf.d
/app/httpd/www
/app/log/httpd /app/log/php 



docker run -dti -v /app/redis/redis.conf:/usr/local/etc/redis.conf --hostname=srv_redis --name=srv_redis sigma_redis:1.0
docker run -dti \
-v /app/php/php.ini:/usr/local/etc/php/php.ini \
-v /app/php/php-fpm.conf:/usr/local/etc/php-fpm.conf \
-v /app/php/pool.conf:/usr/local/etc/php-fpm.d/pool.conf \
-v /app/log/php:/var/log/php \
-v /app/httpd/www:/var/www/html \
--hostname=srv_php \
--name=srv_php \
sigma_php:1.0

docker run -dti \
-v /app/httpd/httpd.conf:/usr/local/apache2/conf/httpd.conf \
-v /app/httpd/vhost.conf:/usr/local/apache2/conf/conf.d/vhost.conf \
-v /app/log/httpd:/var/log/httpd \
-v /app/httpd/www:/var/www/html \
--hostname=srv_php \
--name=srv_php \
sigma_php:1.0
