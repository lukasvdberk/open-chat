### About
An open source discord alternative with privacy, efficiency and hack ability in mind.

**This will be a prototype**


Technology used for this project:

- Go as the programming language for the backend
- Fiber as webframework for the go language
- MySQL for storring for storring information
- Svelte for the front-end

### Installation 
Make sure you add a .env file with the following content at the root of the project.
```env
MYSQL_DATABASE=db_name
MYSQL_ROOT_PASSWORD=password
MYSQL_USER=user
MYSQL_PASSWORD=password
IS_LIVE=false
```
To run this project make sure docker and docker-compose is installed.
Simply then run
```bash
docker-compose up --build
```
The build flag is for building and is only needed the first time.
