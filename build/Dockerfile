FROM alpine:3.3
RUN apk --update --no-cache add \
	ca-certificates
ADD bin/linux/kigo-api /bin/kigo-api
ENTRYPOINT ["kigo-api"]