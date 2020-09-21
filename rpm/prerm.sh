# non standard stuff, stop the service asap
systemctl stop tstwbsrv.service

# unregister the service
%systemd_preun tstwbsrv.service
