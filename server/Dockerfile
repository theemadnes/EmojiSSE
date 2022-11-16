FROM cgr.dev/chainguard/node:latest
ENV NODE_ENV=production

#WORKDIR /usr/src/app

#COPY package*.json ./

#RUN npm --unsafe-perm install

#COPY server.js .

COPY --chown=node:node ["package.json", "package-lock.json*", "./"]

RUN npm install --production

COPY --chown=node:node . .

EXPOSE 8080

CMD [ "server.js" ]
