FROM golang:1.20.4-alpine as stage

WORKDIR /work/
COPY go.* ./
RUN go mod download

COPY main.go main.go
COPY conf/ conf/
COPY data/ data/
COPY handlers/ handlers/

RUN go build -o /work/web.svc

FROM alpine:latest

COPY --from=stage /work/web.svc /work/web.svc
COPY ./template/*.html /work/template/
COPY ./config.json /work/config.json

WORKDIR /work/

EXPOSE ${PORT}
ENTRYPOINT ["./web.svc"]
