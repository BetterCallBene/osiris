FROM ubuntu:jammy
RUN mkdir -p /tmp/go
WORKDIR /tmp/go

RUN apt-get update && apt-get install -y \
    wget\
    git

RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz && rm -rf /tmp/go

ENV PATH=$PATH:/usr/local/go/bin
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

RUN go install -v golang.org/x/tools/cmd/goimports@latest && \
    go install -v golang.org/x/tools/gopls@latest

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# ********************************************************
# * Anything else you want to do like clean up goes here *
# ********************************************************

# [Optional] Set the default user. Omit if you want to keep the default as root.
USER $USERNAME
WORKDIR /usr/local/src