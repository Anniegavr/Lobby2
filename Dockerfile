#FROM golang:1.16-alpine
#RUN mkdir /new
#
#WORKDIR /new
#COPY go.sum ./
#RUN go mod download
#COPY *.go ./
#RUN go build -o /dininghall
#EXPOSE 8081
#CMD ["/dininghall"]


FROM docker.io/library/golang:latest
RUN mkdir /build
WORKDIR /build
RUN export GO111MODULE=on
RUN go get github.com/Anniegavr/Lobby/Lobby
RUN cd /build && git clone https://github.com/Anniegavr/Lobby
RUN cd /build/Lobby/Lobby && go build
EXPOSE 8081
ENTRYPOINT "/build/Lobby/Lobby/main"
