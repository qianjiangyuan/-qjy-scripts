# ingress ssl self-sign CA certs

## general steps

```bash
    # 1. key generation
    openssl genrsa -out <your_key>.key 2048

    # 2. sign request with domain
    openssl req -new -key <your_key>.key -out <your_csr>.csr -subj "/CN=<your_domain>"

    # 3. cert generation
    openssl x509 -req -days 365 -in <your_csr>.csr -signkey <your_key>.key -out <your_cert>.crt

    # 4. kubernetes secret generation
    kubectl create secret tls <your_tls_name> --cert <your_cert>.crt --key <your_key>.key

    # ... after that you can use your tls in your ingress resource
```
