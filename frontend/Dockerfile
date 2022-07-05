FROM node:16-alpine

RUN mkdir -p /usr/src/app

ENV PORT 3000

WORKDIR /usr/src/app

COPY package.json /usr/src/app
COPY yarn.lock /usr/src/app
COPY package-lock.json /usr/src/app

RUN yarn install

COPY . /usr/src/app

EXPOSE 3000

ENTRYPOINT ["yarn", "dev"]