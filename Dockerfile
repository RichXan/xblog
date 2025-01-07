FROM node:18-alpine

WORKDIR /docsify

RUN npm config set strict-ssl false && \
    npm config set registry https://registry.npmmirror.com && \
    npm install -g docsify-cli@latest

EXPOSE 3000

ENTRYPOINT ["docsify"]
CMD ["serve", "."]