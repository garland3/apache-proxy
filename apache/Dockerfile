FROM httpd:2.4-alpine

# Install apache2-utils to get `htpasswd` and create a demo user
RUN apk add --no-cache apache2-utils \
	&& htpasswd -bc /usr/local/apache2/conf/.htpasswd demo demopass

COPY httpd.conf /usr/local/apache2/conf/httpd.conf

EXPOSE 80

CMD ["httpd-foreground"]
