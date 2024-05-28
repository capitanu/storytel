# Notes before

- I guess when the task said "the client can do create/update/delete operations on their message", I guessed no authentication implementation was expected. This implementation expects the client_id to be sent as a PATH param. Of course this would not be something that would go in production, proper authentication with headers (bearer for example with token) would be the solution, but again, assumed this was not expected?
- As it was not really part of the assignment, haven't written unit or integration tests for this, instead used Postman to test it, since it is quite a trivial app

# Running the app (assuming UNIX based systems :please:)

1. Clone the git repository and cd into it

```
git clone git@github.com:capitanu/storytel.git && cd storytel/
```

2. Build the docker image (and give it a tag)

```
docker build -t storytel-calin .
```

3. Run the docker image

```
docker run -p 8180:8180 storytel-calin
```

4. Should now be able to curl `localhost:8180` that will be mapped to the docker container :)

# What to query?

The api spec should be available in [here](./api-spec.yaml)

# CLI? Why not, at least a small one

Make sure you're in the repository directory and that the `client.sh` is executable

```
chmod +x client.sh
```

Then you can run the following commands (while the server is running):

## List all messages

`./client.sh list`

## Get a specific message

`./client.sh list <client_id>`

## Delete a specific message

`./client.sh delete <client_id>`

## Create a message

`./client.sh create <client_id> <content>`

## Update a message

`./client.sh update <client_id> <content>`
