FROM amazonlinux:2
# FROM centos/systemd

RUN yum install -y rpm-build desktop-file-utils python3

RUN curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
  | GH=mh-cbon/changelog sh -xe || true

RUN curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
  | GH=mh-cbon/go-bin-rpm sh -xe

# RUN curl -o /usr/bin/systemctl https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/master/files/docker/systemctl3.py \
#     && chmod 0755 /usr/bin/systemctl
