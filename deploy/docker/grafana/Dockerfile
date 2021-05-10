FROM grafana/grafana:latest

USER root
RUN chown -R root:root /etc/grafana && \
chmod -R a+r /etc/grafana && \
chown -R 472:472 /var/lib/grafana && \
chown -R 472:472 /usr/share/grafana

EXPOSE 3000