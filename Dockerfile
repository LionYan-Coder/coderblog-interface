FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                     /app
ADD bin/linux_amd64/lionblog           $WORKDIR/main
ADD manifest/config/config.pro.yaml    $WORKDIR/config.yaml
ADD resource                    $WORKDIR/resource
RUN chmod +x $WORKDIR/main
###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ["./main","--gf.gcfg.file=./config.yaml"]
