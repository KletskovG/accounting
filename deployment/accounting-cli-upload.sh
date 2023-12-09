date="$(date)"
hash=$(echo $date | sha256sum | cut -d " " -f 1 )
aws s3 --endpoint-url=https://storage.yandexcloud.net cp ./packages/cli/cli $1$hash
curl https://telegram.kletskovg.tech/done/CLI%20binary%20ready%20$hash