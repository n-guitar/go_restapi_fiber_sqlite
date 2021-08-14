FROM golang:1.16-alpine3.14 AS builder
LABEL maintainer="github.com/n-guitar"
RUN apk update && \
    apk upgrade && \
    apk add --no-cache build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:3.14 AS production
RUN apk update && \
    apk upgrade && \
    apk add --no-cache sqlite
COPY --from=builder /app /app
WORKDIR /app
EXPOSE 3000
CMD [ "./main" ]