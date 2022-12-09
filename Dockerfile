FROM golang:1.18

WORKDIR /app

COPY . /app

# install neccessary deb package
RUN apt-get update \
    && apt-get install -y postgresql-client \
    && rm -rf /var/lib/apt/lists/*

RUN chmod +x wait-for-postgres.sh

# build go up
RUN go mod tidy \
    && go mod download \
    && go build -o telegram_bot_mpei ./cmd/main.go


# TODO: this is not good run app for root.
CMD ["./telegram_bot_mpei"]