# go-slackbot
Boiler plate for a golang slack bot via slack slash commands using apigateway and lambda

## Overview
### Handler function
Receives the slack slack command request.
For long running tasks, executes a worker function with a custom payload.

### Worker fucntion
Executed by handler function for longer running tasks

## Setup workspace

Git clone this repository.
`git clone git@github.com:TomSED/go-slackbot.git`

Run `dep ensure` to pull dependencies

### Deployment & Configuration

1. Create a slack command in your Slack App. Executing the command should send a POST request with payload that looks something like this:
```
{
    "token": "example",
    "team_id": "------",
    "team_domain": "------",
    "channel_id": "------",
    "channel_name": "------",
    "user_id": "------",
    "user_name": "------",
    "command": "/slackbot",
    "text": "test",
    "api_app_id": "------",
    "is_enterprise_install": "------",
    "response_url": "https://hooks.slack.com/commands/example/example/example",
    "trigger_id": "------"
}
```

2. Create a `/.env` file according to `/.env.template`. Use `token` from above slack request for env variable `SLACKBOT_AUTH_TOKEN`: 
```bash
$ export $(grep -v '^#' .env | xargs)
```

3. Deploy
```bash
$ make deploy
```

4. Configure your slack slash command request URL using the deployment output. Remember to append the `/slackbot` path
![screenshot](https://imgur.com/7jMV0Qq)