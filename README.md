# React + Golang boilerplate

Note: Optimized for running Docker in Linux where you can use port 80 without running docker-compose as root unless you messed up during docker installation.

## Start
`docker-compose up`  
Visit [localhost](http://localhost)

## Code Away
Some notes while you're coding:
* Changes in ui files will rebuild react app, just need to reload your browser.
* Changes in golang backend app will rebuild and restart the golang app in the docker-compose cluster, you don't even need to reload your browser.
* Changes in the nginx proxy config, the docker-compose file or the Dockerfiles ui/deps.Dockerfile and backend/air.Dockerfile will require a down/up of docker-compose.  
  Either  
  `ctrl-c, docker-compose up`  
  or  
  `docker-compose down && docker-compose up -d`  
  depending on how you docker-compose your stuff.

## Future Improvements
* Use golang endpoint to serve frontend instead of an nginx proxy, to reduce components.