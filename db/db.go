package db

const MongoDBNameEnvName = "MONGO_DB_NAME"

type Pagination struct {
	Limit int64
	Page  int64
}

type Store struct {
	User      UserStore
	Zodiac    ZodiacSignStore
	Horoscope HoroscopeStore
}
