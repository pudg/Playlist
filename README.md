# Playlist
Share short content playlists with close friends or relatives.

# Install
Clone Playlist repository and install required packages.

	cd Playlist
	go mod tidy

Launch the server.

	go run .

Visit [http://0.0.0.0:9000/playlists](http://0.0.0.0:9000/playlists).

See API [API Documentation](#api-documentation) for allowed CRUD operations.

# API Documentation
View Postman generated API documentation here [postman/playlist](https://documenter.getpostman.com/view/29003440/2s9YC8xXL8).


# TODO
- [X] Add skeleton code for REST API
- [X] Add API endpoints documentation
- [X] Add Playlist API Endpoints
  - [X] Add Create playlist functionality (POST)
  - [X] Add Read All/Detail playlist functionality (GET)
  - [X] Add Update playlist functionality (PUT)
  - [X] Add Delete playlist functionality (DELETE)
- [X] Add database connectivity with API
- [ ] Add API test harness
- [ ] Add pagination
- [ ] Add filtering
- [ ] Deploy on AWS