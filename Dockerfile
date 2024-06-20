FROM golang:1.22.3-bookworm

WORKDIR /usr/src/app

RUN apt-get update -y && apt-get dist-upgrade -y && apt-get upgrade -y

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/fallback-control-plane .

ENTRYPOINT ["fallback-control-plane", "-debug"]