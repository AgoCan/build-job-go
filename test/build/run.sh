#!/bin/sh
echo hankbook-web:v1 \| hostname: `hostname` > /var/www/hostname.html
httpd -p 80 -h /var/www -f