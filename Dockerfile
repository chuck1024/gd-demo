FROM registry.cn-chengdu.aliyuncs.com/godog/alpine:3.8

LABEL maintainer="chuck.ch1024@outlook.com"

# 环境变量设置
ENV APP_NAME gd-demo
ENV APP_ROOT /var/www
ENV APP_PATH $APP_ROOT/$APP_NAME
ENV LOG_PATH $APP_ROOT/$APP_NAME/log

# 执行入口文件添加
WORKDIR $APP_PATH
RUN mkdir conf
ADD conf $APP_PATH/conf
RUN mkdir docs
ADD docs $APP_PATH/docs
ADD ./$APP_NAME $APP_PATH/
ADD ./docker/*.sh $APP_PATH/
RUN chmod +x $APP_PATH/*.sh
