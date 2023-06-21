# Docker_Studies

## Docker Essencials
### Interruption
Stopping the container:
```
Docker stop container_id
```

### Listing
Listing containers that are running:
```
docker container ls
```
Listing containers that are not running anymore:
```
docker container ls -a
```
Listing images:
```
docker images ls
```

### Removing
Removing a container:
```
Docker rm container_id
```
Removing an image:
```
docker rmi image_name
```

### Naming
Giving a name to a container
```
docker run -dti --name desired_container_name image_name
```

### Running Configurations
Runninng iteractive container:
```
docker run -it container_id
```
Running detached container 
```
docker run -d container_id
```
Running a command inside a conatainer 
```
docker exec -it container_id command
```

### Moving Arquives
Creating a directory inside a container:
```
docker exec container_id mkdir /destination
```
Listing a directory inside a container
```
docker exec container_name ls /
```
Coping a arquive to the created directory
```
docker cp Arquive.txt container_id:/destination
```
Coping an arquive from the container directory
```

## Installing Applications in a Linux Container
Must run first:
```
apt update
```
And
```
apt upgrade -y
```
Finally:
```
apt -y install application_name
```

