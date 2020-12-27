FROM golang:1.15

ENV PROTOC_VERSION=3.14.0
ENV EVANS_VERSION=0.9.1

######################
# ESSENTIAL PACKAGES #
######################
RUN apt-get update && \
    apt-get install \
    git
    curl \
    unzip \
    -y

##########
# PROTOC #
##########
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    ls -la && \
    unzip -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local bin/protoc && \
    unzip -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local 'include/*'  && \
    rm -f protoc-${PROTOC_VERSION}-linux-x86_64.zip

#################
# PROTOC GEN GO #
#################
RUN go get -u github.com/golang/protobuf/proto \
    github.com/golang/protobuf/protoc-gen-go


#########
# EVANS #
#########
RUN curl -OL https://github.com/ktr0731/evans/releases/download/${EVANS_VERSION}/evans_linux_amd64.tar.gz && \
    ls -la && \
    tar -xvf evans_linux_amd64.tar.gz && \
    mv evans /usr/bin && \
    rm -f evans_linux_amd64.tar.gz