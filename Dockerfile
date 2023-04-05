FROM golang:1.20.3-bullseye AS build-env

WORKDIR /go/src/github.com/Atrix/Atrix

RUN apt-get update -y
RUN apt-get install git -y

COPY . .

RUN make build

FROM golang:1.20.3-bullseye

RUN apt-get update -y
RUN apt-get install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/Atrix/Atrix/build/Atrixd /usr/bin/Atrixd

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["Atrixd"]
