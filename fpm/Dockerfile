FROM php:7.4.15-fpm

RUN php -r "readfile('http://getcomposer.org/installer');" | php -- --install-dir=/usr/bin/ --filename=composer


COPY ./ /opt/app

WORKDIR /opt/app

RUN apt-get update && apt-get install -y git

RUN apt-get install -y autoconf zlib1g-dev

RUN pecl install grpc
RUN echo extension=grpc.so >> /usr/local/etc/php/php.ini

RUN pecl install protobuf
RUN echo extension=protobuf.so >> /usr/local/etc/php/php.ini


ENTRYPOINT ["php-fpm"]