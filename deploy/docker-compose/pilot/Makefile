# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

default: ca pilot-server pilot-client

ca: openpitrix-ca.key openpitrix-ca.crt
pilot-server: pilot-server.key pilot-server.crt
pilot-client: pilot-client.key pilot-client.crt

# openpitrix ca
openpitrix-ca.key:
	openssl genrsa -out openpitrix-ca.key 2048

openpitrix-ca.crt: openpitrix-ca.key
	openssl req -new -x509 -days 3650 \
		-subj "/C=GB/L=China/O=openpitrix/CN=openpitrix.io" \
		-key openpitrix-ca.key -out openpitrix-ca.crt

# pilot server
pilot-server.key:
	openssl genrsa -out pilot-server.key 2048

pilot-server.crt: pilot-server.key openpitrix-ca.key openpitrix-ca.crt
	openssl req -new \
		-subj "/C=GB/L=China/O=pilot-server/CN=pilot.openpitrix.io" \
		-key pilot-server.key \
		-out pilot-server.csr
	openssl x509 -req -sha256 \
		-CA openpitrix-ca.crt -CAkey openpitrix-ca.key -CAcreateserial -days 3650 \
		-in pilot-server.csr \
		-out pilot-server.crt

# pilot client
pilot-client.key:
	openssl genrsa -out pilot-client.key 2048

pilot-client.crt: pilot-client.key openpitrix-ca.key openpitrix-ca.crt
	openssl req -new \
		-subj "/C=GB/L=China/O=pilot-client/CN=pilot.openpitrix.io" \
		-key pilot-client.key \
		-out pilot-client.csr
	openssl x509 -req -sha256 \
		-CA openpitrix-ca.crt -CAkey openpitrix-ca.key -CAcreateserial -days 3650 \
		-in pilot-client.csr \
		-out pilot-client.crt

clean:
	-rm *.key *.crt *.csr *.srl
