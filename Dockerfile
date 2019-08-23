FROM golang

RUN go get github.com/gorilla/mux

RUN go get -v github.com/machinebox/sdk-go/classificationbox

WORKDIR /go/src/github.com/heaptracetechnology/machinebox-classificationbox

ADD . /go/src/github.com/heaptracetechnology/machinebox-classificationbox

RUN go install github.com/heaptracetechnology/machinebox-classificationbox

ENTRYPOINT ["machinebox-classificationbox"]

EXPOSE 8000