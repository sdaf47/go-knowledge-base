openssl req -nodes -new -days 365 -newkey rsa:1024 -x509 -keyout serverkey.pem -out servercert.pem -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=example.com"
cat serverkey.pem servercert.pem >> /etc/stunnel/stunnel.pem