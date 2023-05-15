# goredidefaultstarter
 go redis default REST API with starter Script



Makefile is having :  building, running, stopping, and viewing the logs of your Docker Compose services. It also includes a test target for sending HTTP requests to your service. Here are the descriptions of each target:

- `build` target: This target will build your Docker images using the `docker-compose build` command.
- `run` target: This target will start your Docker Compose services in detached mode using the `docker-compose up -d` command.
- `stop` target: This target will stop your running Docker Compose services using the `docker-compose down` command.
- `logs` target: This target will tail the logs of your Docker Compose services using the `docker-compose logs -f` command.
- `test` target: This target will send HTTP GET and POST requests to your service and display the responses.

You can run these commands using the `make` command followed by the target name. For example, `make build` to build your Docker images, `make run` to start your services, etc. 

Please note that these commands should be run in the directory that contains your `Makefile`, `Dockerfile`, and `docker-compose.yml`.

The `.PHONY` target is used to specify that the targets are not associated with files. This means that `make` will run the commands associated with the targets even if there are files with the same names as the targets in the directory.

Just make sure that the paths and ports in the `test` target match the paths and ports that your service is actually using.