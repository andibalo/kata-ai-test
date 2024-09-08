
# About

Pokemon BE for kata.ai test


## Installation

Ensure you have the following installed
- Docker
- docker-compose
- ngrok (for tunneling)

To setup `ngrok` follow this [tutorial](https://ngrok.com/docs/getting-started/?os=windows)

Since this service is not deployed, in order to access this service from kata.ai platform we will need to start the service locally and use `ngrok` for tunneling to enable public access.

After installing `ngrok` run the following command to start up the db and service
```bash
 docker-compose up
```
    
Then run the following command to enable tunneling so your service is publicly accessible
```bash
 ngrok http http://localhost:8081
```

You will receive a url to access your service for example `https://2ecd-180-254-68-187.ngrok-free.app`

Use this URL as value for the key `pokemonBEURL` in kata.ai platform config
## Architecture

You can view the high level architecture [here](https://drive.google.com/file/d/1gDttQu303wf1xGRSR0ruySMLHnUIqb8r/view?usp=sharing) 
## API Reference

#### Register User

```http
  POST /api/v1/user/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required** |
| `email`      | `string` | **Required** |

#### Login

```http
  POST /api/v1/user/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | **Required** |




## Demo

https://drive.google.com/file/d/1fRtK8-3kVttMkq4lzx8rmgj-1RxEZumj/view?usp=sharing

