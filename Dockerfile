FROM golang:1.13 AS build

COPY ./ /go/src/mando/
WORKDIR /go/src/mando/
RUN go install

FROM cloudfoundry/run:tiny
COPY --from=build /go/bin/mando /mando/
WORKDIR /mando
EXPOSE 8080
ENTRYPOINT ["/mando/mando"]
