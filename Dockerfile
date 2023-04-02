FROM golang:1.20.2-windowsservercore-ltsc2022

WORKDIR /app

COPY Windows .
COPY *.go .
COPY go.mod .
COPY go.sum .

RUN set GOPROXY=direct
RUN go mod download
RUN go build

CMD [ "SteamCMD.exe" ]
