# Base image
FROM centos:7.4.1708

# update yum / timezone / install utils
RUN yum update -y \
  && cp -p /usr/share/zoneinfo/Japan /etc/localtime \
  && yum install -y \
    curl \
    less \
    vim \
    diffutils \
  && yum clean all \
;

#RUN yum-config-manager --add-repo https://getenvoy.io/linux/centos/tetrate-getenvoy.repo
#
#RUN yum-config-manager --enable tetrate-getenvoy-nightly
#
#RUN yum install -y getenvoy-envoy

COPY ./centos7-envoy/nginx_repo/nginx.repo /etc/yum.repos.d/nginx.repo

RUN yum install -y nginx

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
