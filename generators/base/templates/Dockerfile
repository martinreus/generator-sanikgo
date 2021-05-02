# Start by building the application.
FROM golang:1.13-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -mod vendor -o /go/bin/app ./cmd/main.go


# Now copy it into our base image.
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
# copy api files into container
COPY ./api ./api

CMD ["/app"]