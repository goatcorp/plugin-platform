FROM golang:1.19-bullseye as base

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nothing" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 65532 \
    go-user
RUN mkdir /pb_data && \
    chown go-user:go-user /pb_data

WORKDIR $GOPATH/src/plugin-platform/server/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /main ./cmd/platform/main.go

USER go-user:go-user

ENTRYPOINT [ "/main" ]

FROM gcr.io/distroless/static-debian11

COPY --from=base /main .
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base /pb_data /pb_data

USER go-user:go-user

ENTRYPOINT [ "/main" ]
CMD [ "serve" ]