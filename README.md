# Filereader HTTP server

Simple HTTP server which reads a file from a query param
and returns its content. Tool for debugging purpose.

You can use env var FILEREADER_CHROOT (default to .) which defines the base folder where file will be read.

Use the "file" query parameter to specify the file to read

## Run

FILEREADER_CHROOT=/my_data make run

Navigate to http://localhost:8080/?file=my/wanted/file

## Using the Docker image

docker build -t filereader .
docker run --rm -e FILEREADER_CHROOT=/app -v $PWD:/app -p8080:8080 filereader:latest