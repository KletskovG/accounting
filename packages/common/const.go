package common

const DateLayout = "2006-01-02"
const (
	Second = 1000
	Minute = 60 * Second
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 365 * Day
)

const (
	HeaderContentType = "Content-Type"
	ContentTypeJson   = "application/json"
)

const CliUserConfigPath = "./config.env"

const MODE_PROD = "prod"

type hosts struct {
	TelegramAPIURL string
	StorageAPIURL  string
}

func Hosts() hosts {
	return hosts{
		TelegramAPIURL: "https://telegram.kletskovg.tech",
		StorageAPIURL:  "https://storage.yandexcloud.net",
	}
}
