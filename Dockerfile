FROM jenkins/jenkins

USER root
RUN curl -LO https://go.dev/dl/go1.24.0.linux-amd64.tar.gz \
  && rm -rf /usr/local/go \
  && tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz \
  && ln -s /usr/local/go/bin/go /usr/local/bin/go

ENV PATH="/usr/local/go/bin:$PATH"

RUN go install github.com/boumenot/gocover-cobertura@latest

USER jenkins

EXPOSE 8080