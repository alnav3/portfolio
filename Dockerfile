ARG APP_NAME=main

# Build stage
FROM golang:1.21.2 as build
ARG APP_NAME
ENV APP_NAME=$APP_NAME
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /$APP_NAME

# Production stage
FROM alpine:latest as production
ARG APP_NAME
ENV APP_NAME=$APP_NAME
WORKDIR /root/
COPY --from=build /app/$APP_NAME ./
COPY --from=build /app/style ./style
CMD ./$APP_NAME

