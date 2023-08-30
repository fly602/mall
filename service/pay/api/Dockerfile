FROM go1.17

WORKDIR /opt/mall/pay/api

COPY ./etc etc/
COPY payapi bin/

EXPOSE 8003

USER normal

ENTRYPOINT ["bin/payapi", "-f","etc/pay.yaml"]