# Simple database for DnD

Implemented CRUD:
* create
* read
* update
* delete

## Moving the project
You can dump the database:  
```$pg_dump dnd > filename.sql```.

After cloning onto new maching, load database:  
```$psql dnd < filename.sql```,  
then install all the dependencies for our go project:  
go to project folder and run:  
```$go get ./...```



