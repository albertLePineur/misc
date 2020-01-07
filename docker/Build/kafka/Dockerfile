ARG ALPINE_VERSION
ARG JRE_VERSION
ARG KAFKA_VERSION
ARG SCALA_VERSION

FROM alpine:${ALPINE_VERSION}

RUN apk update && apk add bash openjdk${JRE_VERSION}-jre-headless curl && \
    curl http://mirrors.ircam.fr/pub/apache/kafka/${KAFKA_VERSION}/kafka_${SCALA_VERSION}-${KAFKA_VERSION}.tgz --output /opt/kafka.tar.gz && \
    tar zxf /opt/kafka.tar.gz -C /opt/ && \
    mv /opt/kafka_${SCALA_VERSION}-${KAFKA_VERSION} /opt/kafka && \
    chown -R 1001:1002 /opt/kafka && \
    addgroup -g 1002 data && \
    adduser -h /opt/kafka -s /bin/bash -D -u 1001 -G data data && \
    rm -rf /opt/kafka/bin/windows && \
        
EXPOSE 9092
WORKDIR /opt/kafka
USER data
CMD ["/opt/kafka/bin/kafka-server-start.sh","/opt/kafka/config/server.properties"]
