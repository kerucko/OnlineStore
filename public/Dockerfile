FROM node:21-alpine3.17

WORKDIR /app

COPY . .

RUN set -ex && apk --no-cache add sudo \
    && sudo npm install -g http-server