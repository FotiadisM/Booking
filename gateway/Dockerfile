FROM  node:12.16.3-alpine

WORKDIR /app

COPY ["package.json", "yarn.lock", "tsconfig.json", "./"]

RUN yarn install

COPY dist ./dist

EXPOSE 8080

CMD [ "yarn", "start" ]