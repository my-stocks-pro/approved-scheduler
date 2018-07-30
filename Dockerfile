#FROM golang:alpine as builder
#RUN mkdir /build
#ADD . /build/
#WORKDIR /build
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
#FROM scratch
#COPY --from=builder /build/main /app/
#WORKDIR /app
#CMD main


#GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o entrypoint -ldflags '-w -s' main.go
#


FROM scratch

ADD approved-scheduler /

CMD mkdir config

COPY config/approved-scheduler.yaml config/

ENTRYPOINT ["/approved-scheduler"]