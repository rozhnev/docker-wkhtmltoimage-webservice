# docker-wkhtmltoimage-webservice

Inspired by  [docker-wkhtmltopdf-aas ](https://github.com/openlabs/docker-wkhtmltopdf-aas)

wkhtmltopdf in a docker container as a web service.

This image is based on the 
[ubuntu:20.04](https://hub.docker.com/_/ubuntu).

## Running the service

Run the container with docker run and binding the ports to the host.
The web service is exposed on port 80 in the container.

```sh
docker run -d -P docker-wkhtmltoimage-webservice
```

The container now runs as a daemon.

## Using the webservice

There are multiple ways to generate a PDF of HTML using the
service.

### Uploading a HTML file

This is a convenient way to use the service from command line
utilities like curl.

```sh
curl -X GET -vv http://<docker-host>:<port>/?source=www.example.com
```

where:

* docker-host is the hostname or address of the docker host running the container
## TODO

* 

## Bugs and questions


## Authors and Contributors

## Professional Support
