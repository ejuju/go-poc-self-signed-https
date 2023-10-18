# POC for Go HTTPS server with self-signed certificate

Requirements
- [x] Works locally
- [ ] Works online

### Generate CA's private key and root certificate

```sh
openssl req -x509 \
    -sha256 -days 356 \
    -nodes \
    -newkey rsa:2048 \
    -subj "/CN=juliensellier.com/C=FR/L=Paris" \
    -keyout example.root.key -out example.root.crt
```

### Generate server private key and certificate signing request (CSR)

```
openssl genrsa -out server.key 2048 && \
openssl req -new -key server.key -out server.csr -config server.csr.conf
```

### Generate server certificate

```
openssl x509 -req \
    -in server.csr \
    -CA example.root.crt -CAkey example.root.key \
    -CAcreateserial -out server.crt \
    -days 365 \
    -sha256 -extfile server.cert.conf
```

### Optional: Install client certificate

For Chrome:
1. Go to [chrome://settings/certificates](chrome://settings/certificates)
2. Go to the "Authorities" tab
3. Click "Import" and select the root CA certificate


For Linux:
```sh
sudo cp example.root.crt /usr/local/share/ca-certificates/example.crt
sudo update-ca-certificates
```
