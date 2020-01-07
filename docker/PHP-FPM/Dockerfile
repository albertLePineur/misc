FROM php:fpm

#COPY php.ini /root/php.ini
#COPY pool.conf /root/pool.conf
COPY entrypoint.sh /root/entrypoint.sh

#### Copie des fichiers de configurations
RUN chmod +x /root/entrypoint.sh && \
usermod --uid 1001 www-data && \
groupmod --gid 1001 www-data && \
rm -f /usr/local/etc/php/php-fpm.conf /usr/local/etc/php-fpm.d/* /usr/local/etc/php/php.ini* && \
mkdir -p /var/log/php && \
chown www-data:www-data /var/log/php && \ 
#apt-get update && \
#### Installation des modules supplementaires PHP via le PHPize embarque (docker-php-ext-install)
# Modules sans dependances
#docker-php-ext-install opcache bcmath mysqli pdo_mysql calendar gettext sockets && \
# PGSQL
#apt-get install -y libpq-dev && \
#docker-php-ext-install pgsql pdo_pgsql && \
# GD
#apt-get install -y libpng-dev && \
#docker-php-ext-install gd && \
# BZ2
#apt-get install -y libbz2-dev && \
#docker-php-ext-install bz2 && \
# Intl
#apt-get install -y libicu-dev && \
#docker-php-ext-install intl && \
# Zip
#apt-get install -y libzip-dev && \
#docker-php-ext-install zip && \
# LDAP
#apt-get install -y libldap2-dev && \
#docker-php-ext-install ldap && \
# Soap
#apt-get install -y libxml2-dev && \
#docker-php-ext-install soap && \
## => D'autrs extensions sont disponibles mais leur interet reste tres limite ou specifique
#### Installation extensions via PECL
# Redis
#pecl config-set http_proxy http://tcoueffe:Problemchild1997@proxy-pc.na01.groupesigma.fr:8080 && \
pear config-set http_proxy http://tcoueffe:Problemchild1997@proxy-pc.na01.groupesigma.fr:8080 && \
printf "\n" | pecl install redis && \
docker-php-ext-enable redis

# Installation Composer
#apt-get install -y wget && \
#wget -q -O- composer-setup.php https://getcomposer.org/installer | php -- --install-dir /usr/local/bin --filename=composer

EXPOSE 9000
ENTRYPOINT /root/entrypoint.sh
CMD [ "bash" ]
