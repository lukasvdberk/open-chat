### About
An open source discord alternative with privacy, efficiency and hack ability in mind.

**This will be a prototype**


Technology used for this project:

- Go as the programming language for the backend
- Fiber as webframework for the go language
- MySQL for storing information
- Svelte for the front-end
- Docker with docker-compose to combine everything

### Installation 
Clone the repo to the following location
```
go/src/github.com/lukasvdberk
```
Else you won't get any IDE support.

Make sure you add a .env file with the following content at the root of the project.
```env
MYSQL_DATABASE=db_name
MYSQL_ROOT_PASSWORD=password
MYSQL_USER=user
MYSQL_PASSWORD=password
JWT_SECRET=secret_key
IS_LIVE=false
WEB_PUSH_PUBLIC_KEY=key
WEB_PUSH_PRIVATE_KEY=key
WEB_PUSH_EMAIL=email entered on vapidkeys.com
```

The web push keys are required for web notifications. These need to be of the VAPID spec.

A tool to easily generate such keys like https://vapidkeys.com/

The above settings are only for the backend. After that you need to add a settings.js in 
front-end/src/settings.js with the following content.
```js
export const BASE_API_ENDPOINT="/api/"
```
If you host your api (the backend of this project) at a different url then change the url.

To run this project make sure docker and docker-compose is installed.
Simply then run:
```bash
docker-compose up --build --force-recreate
```
The build flag is for building and is only needed the first time (or after a configuration change). After that simply run:
```bash
docker-compose up
```