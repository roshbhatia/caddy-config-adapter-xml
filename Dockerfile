FROM caddy:2.7.5-builder as builder

WORKDIR /app

COPY . .

RUN xcaddy build --with github.com/roshbhatia/caddy-config-adapter-xml=./ --output /usr/bin/caddy

FROM caddy:2.7.5
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
