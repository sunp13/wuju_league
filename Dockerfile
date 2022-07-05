FROM centos:7

ENV TZ=Asia/Shanghai

COPY ./conf /app/conf
COPY ./logs /app/logs
COPY ./wuju_league  /app

WORKDIR /app

CMD [ "./wuju_league" ]
