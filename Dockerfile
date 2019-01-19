FROM golang:1.11.4
ARG SERVICE_NAME

RUN test -n ${SERVICE_NAME}

ENV PORT=8080
ENV GO111MODULE=on
ENV SERVICE_NAME=${SERVICE_NAME}

WORKDIR /go/src/kafka-demo

# Cache step
COPY go.mod .
COPY go.sum .

# Dependencies
RUN go mod download

# Copy source
COPY . /go/src/kafka-demo

RUN go install ./services/${SERVICE_NAME}

CMD ${SERVICE_NAME}

EXPOSE ${PORT}
