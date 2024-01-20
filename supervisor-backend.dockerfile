# ---------------------
# build supervisor-backend
FROM golang:1.20-bullseye as builder-for-supervisor-backend

## set env
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## build
WORKDIR /opt/illa/illa-supervisor-backend
RUN cd  /opt/illa/illa-supervisor-backend
RUN ls -alh

COPY ./ ./

RUN cat ./Makefile

RUN make all

RUN ls -alh ./bin/illa-supervisor-backend


# -------------------
# build runner images
FROM alpine:latest as runner

WORKDIR /opt/illa/illa-supervisor-backend/bin/

## copy backend bin
COPY --from=builder-for-supervisor-backend /opt/illa/illa-supervisor-backend/bin/illa-supervisor-backend /opt/illa/illa-supervisor-backend/bin/


RUN ls -alh /opt/illa/illa-supervisor-backend/bin/



# run
EXPOSE 8003
CMD ["/bin/sh", "-c", "/opt/illa/illa-supervisor-backend/bin/illa-supervisor-backend"]
