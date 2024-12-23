# Traefik Labs

Notes from [udemy course](https://www.udemy.com/course/the-complete-traefik-training-course/) extended with my own code.

## Configuration

At startup, traefik searches for static configuration in a file named `traefik.yml` in `/etc/traefik/`, or `$XDG_CONFIG_HOME/`, or `$HOME/.config/`, or `. (the working directory)`.

## Trusted TLS certificates

The default TLS certificate generated by traefik works, but is not convenient to use, because browsers display `NET::ERR_CERT_AUTHORITY_INVALID` error due to self-signed nature of the certificate. You can work around this error by clicking on the `continue to unsafe ...` link on the page (or typing `thisisunsafe` in chrome), but it is not a good developer experience.

To resolve this issue permanently, you must obtain a TLS certificate signed by a trusted certificate authority. Trusted certificate authority can be installed using [mkcert](https://github.com/FiloSottile/mkcert) tool.

```sh
# Install required linux binaries.
sudo apt install libnss3-tools

# Install mkcert itself.
brew install mkcert

# Install trusted certificate authority.
mkcert -install
```

Restart your browser to reload the trust store.

You can now generate a TLS certificate that will be signed by the generated certificate authority and treated by browsers as a valid TLS certificate.

```sh
mkcert -cert-file tls-certificate.pem -key-file tls-key.pem *.docker.localhost
```

More detailed instructions on configuring docker compose and traefik can be found [here](https://www.putzisan.com/articles/https-setup-with-traefik-docker-compose-for-local-dev).
