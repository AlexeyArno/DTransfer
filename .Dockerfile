FROM alpine:3.5

ARG app_env
ENV APP_ENV $app_env

RUN apk add --no-cache go


RUN apk add --no-cache git


ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

COPY ./ /usr/lib/go/src/github.com/AlexeyArno/golang-files-transfer
WORKDIR /usr/lib/go/src/github.com/AlexeyArno/golang-files-transfer

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8888