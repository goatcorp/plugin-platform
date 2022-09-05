FROM node:18.8-alpine3.15 as base

# building
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn

COPY . .

RUN yarn build

# running
FROM node:18.8-bullseye-slim

WORKDIR /app

COPY --from=base /app/build build
COPY package.json .

EXPOSE 3000
CMD ["node", "./build"]