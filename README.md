Tinode4Chat - Real-Time Chat Backend with Tinode and MongoDB

Overview

Tinode4Chat is a real-time chat backend that leverages the Tinode Instant Messaging Server for handling chat room messaging and MongoDB for storing user, chat room, and message data. The application allows users to create accounts, join chat rooms, send and receive real-time messages, and handle user presence.

The backend is built using Go, MongoDB as the primary database, and Tinode for WebSocket communication. The application is containerized using Docker and orchestrated with docker-compose.

Features

User Management

	•	Sign up with email and password.
	•	Log in and receive a JWT token for authentication.
	•	View user profile.

Chat Rooms

	•	List available chat rooms.
	•	Create new chat rooms.
	•	Join/leave chat rooms.

Messaging

	•	Send messages to a chat room.
	•	Receive real-time messages in a chat room.
	•	Track presence (online/offline) of users.

Database

	•	MongoDB for persisting user, chat room, and message information.

Running the Application
Set up environment variables:
Create a .tmp/env file with the following content:
```
MONGO_URL=mongodb://root:example@mongodb:27017/tinode?authSource=admin&replicaSet=rs0
TINODE_URL=ws://tinode:6060
TINODE_ADMIN_EMAIL=admin@example.com
TINODE_ADMIN_PASSWORD=your_admin_password
```

Run the application using Docker Compose:
```
scripts/docker-compose.sh
```


Testing the API:
Example API requests can be made using Postman or curl. You can interact with the following endpoints:
- `POST /signup`:
```
curl -X POST -H "Content-Type: application/json" -d '{"email":"user@example.com","username":"user1","password":"password"}' http://localhost:8080/signup
```

- `POST /login`:
```
 curl -X POST -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"password"}' http://localhost:8080/login
```

- `GET /profile`:
```
curl -X GET http://localhost:8080/profile \
  -H "Authorization: Bearer <YOUR JWT TOKEN>"
```

- `GET /rooms`:
```
curl -X GET http://localhost:8080/rooms \
  -H "Authorization: Bearer <YOUR JWT TOKEN>"
```

- `POST /rooms`:
```
curl -X POST http://localhost:8080/rooms \
  -H "Content-Type: application/json" -d '{"name":"test"}' \
  -H "Authorization: Bearer <YOUR JWT TOKEN>"
```

- `POST /rooms/:id/join`:
```
curl -X POST http://localhost:8080/rooms/66e21a0d975def7eaeba7782/join \
  -H "Content-Type: application/json" -d '{"name":"test"}' \
  -H "Authorization: Bearer <YOUR JWT TOKEN>"

```

- `POST /rooms/:id/leave`:
```
curl -X POST http://localhost:8080/rooms/66e225298cba1bb41d74ee2a/leave \
  -H "Authorization: Bearer <YOUR JWT TOKEN>"

```

- `POST /rooms/:id/messages`:
```
curl -X POST http://localhost:8080/rooms/66e234947b6584bf8dcb8b3c/messages \  
  -H "Content-Type: application/json" -d '{"content":"Hello, world!"}' \
  -H "Authorization: Bearer <YOUR JWT TOKEN>"
```

- `GET /rooms/:id/messages`:
```
curl -X GET http://localhost:8080/rooms/66e234947b6584bf8dcb8b3c/messages \ 
  -H "Authorization: Bearer <YOUR JWT TOKEN>"
```

License

This project is licensed under the MIT License.
