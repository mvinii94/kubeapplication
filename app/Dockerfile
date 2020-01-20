FROM alpine:latest

RUN apk add --no-cache make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Copy simple app
RUN mkdir -p /app
COPY . /app
WORKDIR /app

# Compile app
RUN GOOS=linux go build -a -installsuffix cgo -o main main.go
RUN chmod +x main

EXPOSE 8080
ENTRYPOINT "./main"