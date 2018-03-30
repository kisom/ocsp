# OCSP responder

This is the skeleton for an OCSP response server.

The server should implement an HTTP interface to some data store that
contains certificate validity information, replying to requests with
the appropriate status.

## Docs

* [RFC 6960 - X.509 Internet PKI OCSP](https://tools.ietf.org/html/rfc6960) (see
  also the `docs/` directory in this project).
* Go's [`pkix`](https://golang.org/pkg/crypto/x509/pkix/) package containing some
  of the structures used in OCSP.
* Go's [`ocsp`](godoc.org/golang.org/x/crypto/ocsp) package containing
  OCSP parsing and serialisation functions.

## Questions

1. What information is needed from a certificate to make answer
   revocation status questions?
2. What data should be stored?
3. How should that data be stored (e.g. what is the backing data store)?

## Signing certs

A CSR (in [cfssl JSON format](https://github.com/cloudflare/cfssl)) is
present in `certs/`.
