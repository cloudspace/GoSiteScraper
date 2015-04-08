FROM gliderlabs/alpine:edge
RUN apk update
RUN apk-install libxml2 ca-certificates
ADD ./scraper /scraper
