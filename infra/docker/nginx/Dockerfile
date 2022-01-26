FROM nginx:1.21.5-alpine

COPY infra/docker/nginx/default.conf.tmpl /etc/nginx/conf.d/default.conf.tmpl

ENV LOCAL_HOST = 127.0.0.1
EXPOSE 80

CMD /bin/sh -c 'sed "s/\${LOCAL_HOST}/${LOCAL_HOST}/" /etc/nginx/conf.d/default.conf/tmpl > /etc/nginx/conf.d/default.conf && nginx -g "daemon off;"'
