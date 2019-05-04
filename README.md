# MOC to IRC [![](https://images.microbadger.com/badges/image/ctfl/moc-irc.svg)](https://hub.docker.com/r/ctfl/moc-irc "DockerHub Image")

Sends MOC Messages to IRC Channel

## Configuration

### SQLite3 Config

```bash
API_ENDPOINT=http://moc
IRC_ADDRESS=chat.freenode.net:6667
IRC_NICK=MySuperDupaTest
IRC_PASSWORD=
IRC_USER=MySuperDupaTest
IRC_FULLNAME=MySuperDupaTest
IRC_CHANNEL=#my-super-dupa-test
```

## Usage

```
moc-irc
```

## Docker Compose

```yaml
version: '3'
services:
  moc:
    image: ctfl/moc
    env_file:
      - .env
    ports:
      - 80:80
  moc-irc:
    image: ctfl/moc-irc
    env_file:
      - .env
```
