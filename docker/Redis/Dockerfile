FROM redis:5
COPY entrypoint.sh /root/entrypoint.sh
RUN chmod +x /root/entrypoint.sh && \
usermod --uid 1001 redis && \
groupmod --gid 1001 redis && \
mkdir -p /var/log/redis && \
chown redis:redis /var/log/redis
EXPOSE 9000
ENTRYPOINT /root/entrypoint.sh
CMD [ "bash" ]
