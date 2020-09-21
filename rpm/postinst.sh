# add a group
/usr/bin/getent group tstwbsrv >/dev/null || /usr/sbin/groupadd -r tstwbsrv
# add an user
/usr/bin/getent passwd tstwbsrv >/dev/null || /usr/sbin/useradd tstwbsrv -g tstwbsrv --system --no-create-home --home /nonexistent

# create certificate
mkdir -p 0755 /etc/pki/tls/certs/tstwbsrv
/etc/pki/tls/certs/make-dummy-cert /etc/pki/tls/certs/tstwbsrv/localhost.crt || true
chmod 644 /etc/pki/tls/certs/tstwbsrv/localhost.crt

# create run time dir
mkdir -p 0755 /var/run/tstwbsrv
chown -R tstwbsrv:tstwbsrv /var/run/tstwbsrv

# register service
%systemd_post tstwbsrv.service

# non standard stuff, start the service asap
systemctl start tstwbsrv.service
