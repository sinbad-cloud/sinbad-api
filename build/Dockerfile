FROM alpine:3.3
RUN apk --no-cache add \
	ca-certificates
ADD bin/linux/sinbad-api /bin/sinbad-api
ENTRYPOINT ["sinbad-api"]