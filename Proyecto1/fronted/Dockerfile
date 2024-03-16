FROM node:19.0.0-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm ci --silent
COPY . ./
RUN NODE_ENV=production npm run build --silent


FROM nginx:stable-alpine
WORKDIR /usr/share/nginx/html
COPY --from=build /app/dist ./
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]