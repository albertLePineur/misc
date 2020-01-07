FROM httpd:2.4
RUN rm -f /usr/local/apache2/conf/httpd.conf && \
usermod --uid 1001 www-data && \
groupmod --gid 1001 www-data && \
mkdir -p /usr/local/apache2/conf/conf.d/ && /var/www/html 
EXPOSE 80
cmd [ "bash" ]

