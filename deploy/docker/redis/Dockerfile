FROM redis:latest

RUN mkdir -p /usr/local/etc/redis
COPY redis.conf /usr/local/etc/redis/redis.conf

ARG REDIS_PASSWORD=secret
RUN sed -i "s/# requirepass foobared/requirepass ${REDIS_PASSWORD}/g"  /usr/local/etc/redis/redis.conf
RUN sed -i "s/bind 127.0.0.1/requirepass ${REDIS_PASSWORD}/g"  /usr/local/etc/redis/redis.conf

ARG TZ=UTC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

VOLUME /data

EXPOSE 6379

CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]

