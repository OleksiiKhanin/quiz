FROM golang:1.19-alpine3.16 as builder

ENV APP_NAME=cardsvc
WORKDIR /${APP_NAME}

COPY ./api api
COPY ./cmd cmd
COPY ./db db
COPY ./dto dto
COPY ./interfaces interfaces
COPY ./service service
COPY ./go.mod .
COPY ./go.sum .

RUN go build -v -o ${APP_NAME} cmd/main.go

FROM alpine:3.17
COPY --from=builder /cardsvc/cardsvc /bin/
COPY ./migrations/*.sql /migrations/
COPY ./frontend/build/* /static/

CMD ["cardsvc", "-dhost", "postgres", "-duser", "postgres", "-dpass", "postgres", "-m", "/migrations/", "-s", "/static/"]
