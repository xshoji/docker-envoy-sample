# Base image
FROM nginx

# commands
RUN apt-get -y update && apt-get -y upgrade && apt-get -y dist-upgrade 
# install utils
RUN apt-get install -y \
    curl \
    less \
    vim \
;

COPY ./nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
