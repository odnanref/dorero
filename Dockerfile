FROM golang:1.18-bullseye AS build 

COPY app.go .
RUN mkdir templates
RUN go build app.go
COPY ./templates/index.html ./templates/
RUN rm *.go

FROM debian:bullseye

RUN mkdir /templates/
COPY --from=build /go/* /
COPY --from=build /go/templates/* /templates/

EXPOSE 8000
RUN chmod 755 /app
CMD ["/app"]
