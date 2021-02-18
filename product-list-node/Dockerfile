# FROM node:15.8.0
# # ENV NODE_ENV=production

# WORKDIR /opt/app
# COPY package*.json ./

# COPY . .

# RUN yarn

# RUN yarn add tsc

FROM node:15.8.0
WORKDIR /opt/app
COPY package*.json ./

RUN yarn

COPY . .
EXPOSE 3333

CMD ["yarn", "start"]