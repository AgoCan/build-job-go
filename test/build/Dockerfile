FROM busybox:latest
RUN echo hankbook-web:v1 >> /var/www/index.html
COPY run.sh /
RUN chmod +x /run.sh
CMD ["/run.sh"]