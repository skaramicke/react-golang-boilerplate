FROM node:14-alpine
ADD package.json .
ADD yarn.lock .
RUN yarn 
CMD echo "ready"