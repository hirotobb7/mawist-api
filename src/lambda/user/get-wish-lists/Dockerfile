FROM public.ecr.aws/bitnami/golang:1.16 as build-image

WORKDIR /go/src

COPY ./go.mod ./
COPY ./go.sum ./

COPY ./lambda/user/get-wish-lists/main.go ./lambda/user/get-wish-lists/

COPY ./internal/dynamo ./internal/dynamo
COPY ./pkg/json ./pkg/json
COPY ./pkg/response ./pkg/response
COPY ./pkg/validator ./pkg/validator

WORKDIR /go/src/lambda/user/get-wish-lists

RUN go build -o /go/bin

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["get-wish-lists"]
