ARG APP_NAME=app

# Build stage
FROM golang:1.21.2 as build
ARG APP_NAME
ENV APP_NAME=main
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /main .

# Production stage
FROM alpine:latest as production
ARG APP_NAME
ENV APP_NAME=main

# Add a new user "appuser" and switch to it
RUN adduser -D appuser
USER appuser
WORKDIR /home/appuser

COPY --from=build /main ./
COPY --from=build /app/style ./style

# Use an absolute path for CMD
CMD ["/home/appuser/main"]

