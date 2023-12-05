FROM ubuntu:latest

LABEL maintainer="mismardev"

RUN apt-get update && \
    apt-get install -y -q curl gnupg2
RUN curl https://nginx.org/keys/nginx_signing.key | apt-key add -

RUN apt-get update && \
    apt-get install -y -q nginx && \
    apt-get install certbot

EXPOSE 443 80

CMD ["nginx", "-g", "daemon off;"]