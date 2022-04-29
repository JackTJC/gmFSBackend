# from: https://nodeployfriday.com/posts/self-signed-cert/
# 证书生成
generate:
	openssl req -new -nodes -x509 -days 365 -keyout domain.key -out domain.crt -config ./cert.conf -extensions v3_req

validate:
	openssl x509 -in domain.crt -noout -text

cert_2_perm:
	openssl x509 -in domain.crt -out domain.pem
