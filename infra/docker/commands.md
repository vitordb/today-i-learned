### deleting containers

docker container rm 'container id/name'

### creating container and entering

docker container run -it ubuntu /bin/bash

### detached container

docker container run -d nginx

### publishing port

docker container run -d -p 8080:80 nginx
