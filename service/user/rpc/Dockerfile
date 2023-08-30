FROM go1.17

WORKDIR /opt/mall/user/rpc

COPY ./etc etc/
COPY userrpc bin/

USER normal

EXPOSE 9000

ENTRYPOINT ["bin/userrpc", "-f","./etc/user.yaml"]