# ReREST
[![GoDoc](https://godoc.org/github.com/willyrgf/rerest?status.png)](https://godoc.org/github.com/willyrgf/rerest)

Serve a RESTful API from any Redis database

## Redis version

- 3.0 or higher

## Environment
```
REREST_CONF="config.toml"
```

## Usage
```
Usage of ./rerest:
  -dev
        Set the environment to dev.
  -trace
        Enable trace.
```

### Test with docker-compose
```
docker-compose up --build -d
./test_example.sh
```

### For build:
```sh
git clone https://github.com/willyrgf/rerest.git
cd rerest
./build.sh
```

### Install using go:
```sh
go install github.com/willyrgf/rerest
```

### Configure like a daemon in FreeBSD:
```sh
cat <<EOF >> /etc/rc.conf
# ReREST
rerest_enable="YES"
rerest_conf="/usr/local/etc/rerest/config.toml"
EOF
```
