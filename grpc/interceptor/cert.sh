openssl genrsa -out ca-key.pem 2048

openssl req -new -x509 -nodes -days 365 \
   -key ca-key.pem \
   -out ca-cert.pem \
   -subj "/C=CH/CN=localhost" \
   -addext "subjectAltName = DNS:localhost" \