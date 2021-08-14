FROM golang:1.16.4 AS builder
LABEL maintainer="github.com/n-guitar"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:3.14 AS production
RUN apk update && \
    apk upgrade && \
    apk add --no-cache sqlite
RUN mkdir /data
COPY --from=builder /app /app
WORKDIR /app
CMD [ "./main" ]