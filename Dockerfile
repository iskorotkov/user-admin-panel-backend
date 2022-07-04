FROM golang:1.18-alpine AS build
WORKDIR /go/src
ENV CGO_ENABLED=0

COPY ["go.mod", "go.sum", "./"]
RUN go get -d -v ./...

COPY . .
RUN go build -a -installsuffix cgo -o app .

FROM scratch AS runtime
COPY --from=build /go/src/app ./

EXPOSE 8080/tcp
ENTRYPOINT ["./app"]
