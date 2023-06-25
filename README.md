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

### Creating a New Place for Data Storage
Through the command "docker inspect" it is possible to verify that the standart storage is placed inside the container. In order to change this place we do:
```
docker run -e MYSQL_ROOT_PASSWORD=desired_password --name container_name -d -p 3306:3306 --volume=/new_place:old_place image_name
```
This original place can be found in the volume type after running the "docker inspect" command.

### The Data Storage Types
The previous case of mounting is called bind. There are two more types: the named and Dockerfile volume.
  - Bind: bindding a host directory inside the container.
    - ```docker run -v /hostdir:/containerdir image_name```
  - Named: there are ways to create docker volumes. They are created inside a specific directory in the host
    - Creating a Docker Volume: ```docker volume create volume_name```
    - Configuring the container to store data inside the created volume: ```docker run -v volume_name:/containerdir image_name```
  - Dockerfile Volume: volumes that are created by the dockerfile instructions. They are also created inside a specific directory, but they do not have a name.

### Other Storage Commands
It is also possible to create a bind mount by using the "mount" command as well.
```
docker run -dti --mount type=bind,src=/hostdir,dst=/containerdir image_name
```
Another possibility is to create a read only mount:
```
docker run -dti --mount type=bind,src=/hostdir,dst=/containerdir,ro image_name
```
Cheking the volumes that already exists:
```
docker volumes ls
```
Creating a volume:
```
docker volume create volume_name
```
Attaching a container to a volume:
```
docker run -dti --name containe_name --mount type=volume,src=volume_name,dst=/data image_name
```

### Creating a Web Application with Apache
The Apache image is called "httpd".
```
docker pull httpd
```
The content to be displayed is placed inside a directory in the host and this directory is binded to the container.
```
docker run -d --name apache_container -p 80:80 --volume=/hostdir:/usr/local/apache2/htdocs httpd
```

## Controlling Memory and CPU
In order to visualize the memory and processing utilization by a container we use the ```docker stats container_name```.
To change the configuratios of a conatainer we use the "update" command. To update the maximun memory and processing capability that a created container is allowed to use we run:
```
docker update container_name -m 128M --cpus 0.2
``` 

## Other Commands
- ```docker info```: server informations.
- ```docker container logs container_name```: execution logs.
- ```docker container top container_name```: processes in execution.

## Network
Listing the avaiable networks
```docker network ls```
Every created container is added to the bridge network and has acess to the host network, thats the reason that when we indicate the host ip we have acess to the containers.
```docker network inspect bridge``` Shows the containers in the bridge network.
Containers in the same network have acess to each other. If you want to isolate a group of containers, it possible to create a network and place the container in there. Creating a network: ```docker network create network_name```.
Placing a container inside the network created: 
```
docker run -dti --name container_name --network network_name image_name
```
Deleting a network: ```docker network rm network_name```

## Dockerfile
To create an application inside a container without using the Dockerfile, it is necessary to create a container interactive trought the ```run -dti``` command and make the alterations inside the container manually.
  1. ```docker run -dti --name ubuntu-python ubuntu```
  2. ```docker exec -ti ubuntu-python bash```
  3. ```apt update```
  4. ```apt install python3 vim```
  5. ```apt clean```
  6. ```cd app_dir```
  7. ```vi app.py```
  8. Write Code
  9. Outside the container: ```docker exec -ti ubuntu-python python3 /app_dir/app.py```
WHen using a Dockerfile it possible to do all these steps in a file.
```
FROM ubuntu
RUN apt update && apt install -y python3 && apt clean
COPY /hostdirr/app.py /containerdir/app.py
CMD python3 /containerdir/app.py
```
Then we build the dockerfile: ```docker build dockerfile_path -t image_name```

## Project 1: Web Server with Dockerfile
To create a web server we will be using apache installed inside a debian container. All the container configuration will be done trough a docker file.
First we will get the web structure trough a "wget" and, since there is a commanda in dockerfile that discompacts a .tar arquive we will be using such a format.
```
wget
```
Then we will write the dockerfile content in witch we will define the base memory (```FROM```), some environmental variables (```ENV```) and some commands (```CMD```).
