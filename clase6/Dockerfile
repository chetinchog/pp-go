FROM golan:1.14
COPY . /go/src/github.com/chetinchog/pp-go/clase6
WORKDIR /go/src/github.com/chetinchog/pp-go/clase6
RUN go build -o /go/bin/clase6 .

FROM alpine
RUN apk --update add ca-certificates
RUN mkdir -p /app/
COPY --from=0 /go/bin/clase6 /app/clase6
WORKDIR /app
ENTRYPOINT [ "clase6" ]