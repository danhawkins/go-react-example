# Example app with Go/Gin/React/Typescript

An example app that compiles a react typescript app to golang

## Resources
https://www.youtube.com/watch?v=p08c0-99SyU



Add domain/todos folder which will contain the service, the service should be passed a repo to be used, and the repo should be passed a db to be used

Then have a seperate folder for routes / handlers (web stuff) keep the the only thing that should be exposed from the domain is the service layer, then the
web layer will use the service for all interactions, perhaps we have a todos_web folder which will expose the web based stuff specifically. Then when we run main
we will initialise a service and pass it to the web setup.

When we boot main we init the database, the have a setupTodos while will create a new service and mount the web stuff
