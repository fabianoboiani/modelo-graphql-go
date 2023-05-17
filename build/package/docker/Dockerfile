# Start from the latest golang base image
FROM golang:1.18 as builder


ARG OS_NAME
ARG ARCH

ENV GOOS=${OS_NAME:-linux} \
    GOARCH=${ARCH:-amd64}

WORKDIR /code
COPY .  .

RUN echo ${GOOS}/${GOARCH} > scripts/plataforms \
    && make clean config-private-repo deps test build \
    && ls -la target/bin/${GOOS}_${GOARCH}



######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

LABEL maintainer="Lojas Renner DataLab" \
 contributor="Fabiano Boiani <fabiano.boiani@lojasrenner.com.br>"

ARG OS_NAME
ARG ARCH

ENV GOOS=${OS_NAME:-linux} \
 GOARCH=${ARCH:-amd64} \
 APP_NAME=modelo-graphql-go

WORKDIR /app/

COPY --from=builder /code/target/bin/${GOOS}_${GOARCH}/${APP_NAME}_*_${GOOS}_${GOARCH}  ./${APP_NAME}
COPY ./configs/application.yaml ./configs/application.yaml

CMD ./${APP_NAME} -e application