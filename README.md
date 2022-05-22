# URLStalker 

App to scrape URLs at regular interval

## Architecture 

* Resources: API for storing and retrieving URLs that need to be tracked, including any relevant metadata
* Snapshots: API for NoSQL DB with snapshots of URLs 
* Scraper: Go application for scraping the URLs 

![Architecture](docs/architecture.svg)

## Development 

The development environments of the three apps are completely separate and decoupled. The idea is that you can optimize your development environment to the app at hand. 

# To do

- [x] Create Snapshot app 
- [x] Create docker compose file for deployment 
- [ ] Add persistent storage to snapshot app 
- [ ] Add read end-points to snapshot app
- [ ] Create scraper app

