date="$(date)"
hash=$(echo $date | sha256sum | cut -d " " -f 1 )
aws s3 --endpoint-url=https://storage.yandexcloud.net cp ./packages/cli/cli_linux $1$hash_linux
aws s3 --endpoint-url=https://storage.yandexcloud.net cp ./packages/cli/cli_macos_x86 $1$hash_macos_x86
curl https://telegram.kletskovg.tech/done/CLI%20binary%20ready
curl https://telegram.kletskovg.tech/done/$hash