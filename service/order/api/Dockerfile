FROM go1.17

WORKDIR /opt/mall/order/api

COPY ./etc etc/
COPY orderapi bin/

EXPOSE 8002

USER normal

ENTRYPOINT ["bin/orderapi", "-f","etc/order.yaml"]