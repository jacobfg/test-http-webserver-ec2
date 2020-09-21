# not sure what this does...
%systemd_postun_with_restart tstwbsrv.service

# remove certificate
rm -rf /etc/pki/tls/certs/tstwbsrv

rm -rf /var/run/tstwbsrv
