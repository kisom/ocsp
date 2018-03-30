CERTS :=	ca.pem signer.pem server.pem
JSONS :=	ca.json signer.json server.json config.json

all: $(CERTS)

$(CERTS): $(JSONS)
	./setup.sh

.PHONY: clean
clean:
	rm *.pem *.csr
