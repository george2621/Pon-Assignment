# Pon-Assignment
It's an API written in Go to show message with the ability to change the message 
and see the number of people who saw it before 

## Features
1. I Used gorilla/mux laibrary
2. Deploy it by Elastic Beanstalk <a href="http://Golang-api-env.eba-nwbdfmij.us-east-1.elasticbeanstalk.com ">AWS</a>
3. Use HTTPS instad of http 

## CLI command to create cerificate and key for HTTPS protocol
### certificate signing request (CSR):
 This command generates a localhost.key file which is the private key and localhost.csr
 which is the certificate signing request that contains the public key.
> $ openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr

### public key:
This command generates the localhost.crt file which is the self-signed certificate signed 
by our own localhost.key private key. The x509 flag states the standard format of an SSL/TLS certificate which is X.509.
> $ openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt

