FROM node:18.8-alpine3.15 as base

ARG VITE_API_ADDRESS
ARG VITE_PUBLIC_API_ADDRESS
ENV VITE_API_ADDRESS $VITE_API_ADDRESS
ENV VITE_PUBLIC_API_ADDRESS $VITE_PUBLIC_API_ADDRESS

# building
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn

COPY . .

RUN yarn build

# running
FROM node:18.8-bullseye-slim

WORKDIR /app

COPY --from=base /app .
COPY . .

EXPOSE 3000
CMD ["node", "./build"]