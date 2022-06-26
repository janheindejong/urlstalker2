package main

func main() {

	config := LoadConfigFromEnv()
	dbService := DbService{host: config.ApiHost}
	emailService := EmailServiceSmtp{
		host:    config.EmailHost,
		port:    config.EmailPort,
		account: config.EmailAccount,
		pass:    config.EmailPassword,
	}

	scraper := Scraper{
		dbService:    &dbService,
		emailService: &emailService,
		waitDuration: config.WaitDuration,
	}

	scraper.RunForever()

}
