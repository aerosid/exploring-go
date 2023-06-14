FROM ubuntu:22.04
ARG UID
ARG GID
RUN addgroup --gid $GID ubuntu \
&& adduser --disabled-password --gecos "" --gid $GID --uid $UID ubuntu \
&& apt update \
&& apt install -y vim dnsutils openssl curl wget tree file git \
&& wget -O /tmp/go1.20.5.linux-amd64.tar.gz https://go.dev/dl/go1.20.5.linux-amd64.tar.gz \
&& rm -rf /usr/local/go  \
&& tar -C /usr/local -xzf /tmp/go1.20.5.linux-amd64.tar.gz \
&& rm /tmp/go1.20.5.linux-amd64.tar.gz
COPY .bash_profile /home/ubuntu/
# See: https://www.baeldung.com/linux/bashrc-vs-bash-profile-vs-profile
USER ubuntu:ubuntu
WORKDIR /home/ubuntu
