FROM go1.17

WORKDIR /opt/mall/product/api

COPY ./etc etc/
COPY productapi bin/

EXPOSE 8001

USER normal

ENTRYPOINT ["bin/productapi", "-f","etc/product.yaml"]