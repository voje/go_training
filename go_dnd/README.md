# Simple database for DnD

Implemented CRUD:
* create
* read
* update
* delete

## Moving the project
You can dump the database:  
```$pg_dump dnd > filename.sql```.

After cloning onto new maching, you can restore the db there:
* first, create a new database:  ```CREATE DATABASE dnd```,  
* then restore from file: ```$psql dnd < filename.sql```,  
then install all the dependencies for our go project  
by going to project folder and running:  
```$go get ./...```



