# Docker_Studies

## Docker Essencials
### Interruption
Stopping the container:
```
Docker stop container_id
```

### Restarting a Container
To restart a stopped container we run:
```
docker start container_id
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
docker cp container_name:/destination/Arquive.txt my_computer
```

### Tags
The tags are related to the image's versionsand are placed after the name of the image
```
docker pull image_name:image_tag
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

## Creating a MySQL Container
The image name is "mysql", and, beyond that, in this case it is necessary to specify some enviroment variables such as the root password.

### Pulling the MySQL Image
To pull the image we use the command above.
```
docker pull mysql
```

### Creating the Cotainer with the Environment Variables
To especify the environment variables we use the "-e" sign.
```
docker run -e MYSQL_ROOT_PASSWORD=desired_password --name container_name -d -p 3306:3306 image_name
```
In the previous command we also specify the port that will be heard and the port to obtains the information.

### Creating a Database Inside the Container
To enter the container we run:
```
docker exec -it container_name bash
```
Already inside the container, in order to create a database, we log in the MySQL service: 
```
mysql -u root -p --protocol=tcp
```
And then create the table
```
CREATE DATABASE database_name;
```
To view the available databases:
```
show databases;
```

