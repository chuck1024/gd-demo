FROM loads/alpine:3.8

LABEL maintainer="chuck.ch1024@outlook.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /home/work/app/gd-demos

# 添加应用可执行文件，并设置执行权限
ADD ./bin/linux_amd64/main   $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加配置文件
ADD conf   $WORKDIR/conf

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
