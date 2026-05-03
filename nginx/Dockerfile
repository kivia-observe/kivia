FROM nginx:trixie-perl

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
EXPOSE 443

CMD ["/bin/sh", "-c", "while sleep 6h; do nginx -s reload; done & nginx -g 'daemon off;'"]
