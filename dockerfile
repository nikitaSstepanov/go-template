FROM golang:alpine AS builder

EXPOSE 80

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]

RUN ["go", "mod", "download"]

COPY ./ ./

RUN ["go", "build", "-o", "./bin/app", "./cmd/app"]

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app ./

COPY ./config ./config

COPY ./migrations ./migrations

COPY .env .env

CMD [ "./app" ]