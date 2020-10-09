FROM golang:1.15 AS build

COPY ./ /go/src/mando/
WORKDIR /go/src/mando/
RUN go install

FROM gcr.io/distroless/base:nonroot
COPY --from=build /go/bin/mando /mando/
WORKDIR /mando
EXPOSE 8080
USER 65532
ENTRYPOINT ["/mando/mando"]
