FROM golang:1.17-bullseye AS build

WORKDIR /src

COPY . /src/

RUN make


FROM gcr.io/distroless/base

COPY --from=build /src/bin/kohaku /app/
COPY ./conf/config.yaml /

CMD ["./app/kohaku"]
