# Projects

## al2 projects list
This Command is used to : 

	- Get all projects.

#### Command

```shell
al2 projects list --orgunitID=ORG_ID  [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 projects list --orgunitID=1 # Gets all project having having orgunitID = 1 
al2 projects list --orgunitID=1 --limit=10 # Gets 10 project having having orgunitID = 1
al2 projects list --orgunitID=1 --sort-by="name" # Gets all project order by name having having orgunitID = 1
al2 projects list --orgunitID=1 --sort-by="created_at" --limit=10 # Gets the 10 first project ordred by created date having orgunitID = 1
al2 projects list --orgunitID=1 --sort-by="~created_at" --limit=10 # Gets the 10 last projects ordred by created date desc having orgunitID = 1
al2 projects list --orgunitID=1 --filter="id,name" # Gets IDs, names of all projects having orgunitID = 1
al2 projects list --orgunitID=1 --filter="id" # Gets IDs of all projects having orgunitID = 1
```

#### Output
```shell
| ID                                   | Name        | Description |
| ------------------------------------ | ----------- | ----------- |
| 22b007c9-c6c6-4100-8e77-6a667c7df856 | project1    | description |
| f4ef44f4-7dbd-4648-b782-f9df29b7d40c | project2    | description |
```


## al2 projects get
This Command is used to : 

	- Get a project by their ID.

#### Command

```shell
al2 projects get --id=PROJECT_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 projects get --id=1 # get project having ID = 1
```

#### Output
```shell
	Name : project1
	Description : description project 1
	OrgUnitID : 123
```
#### 


## al2 projects create
This Command is used to : 

	- Creates a project.

#### Command

* `al2 projects create -f FILEPATH ` - Create project using file. ( json / yaml)
* `al2 projects create --name=NAME --description=DESCRIPTION --orgunitID=ORG_ID ` - Create project using flags.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 projects create --name="project 3" --description="description 3" --orgunitID=1 # Create project using flags.
al2 projects create -f "C://user/..../file.json" # create project using json file.
al2 projects create # Create project using system Interaction.
```

#### Output
```shell
	OK / KO
```
#### 

## al2 projects update
This Command is used to : 

	- Updates a project.

#### Command

* `al2 projects update --id=PROJECT_ID -f FILEPATH ` - Update project using file. ( json / yaml)

* `al2 projects update --id=PROJECT_ID --name=NAME --description=DESCRIPTION --orgunitID=ORG_ID` - Update project using flags.

* `al2 projects update --id=PROJECT_ID` - Update project using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to update the resource.

#### Examples

```shell
al2 projects update --id=1 --name "CLI" --description "commands line interface" # update name and description of the project having id = 1
al2 projects update --id=1 --name "CLI" # update name of the project having id = 1
al2 projects update --id=1 --description "commands line interface" # update description of the project having id = 1
al2 projects update --id=1 -f "C://user/..../file.json" # update project using json file.

#interction with system
al2 projects update --id=1
	New project name ("project1") :
	New project description ("description"): Command line interface
	New project OrgunitID ("12") :
#in this case we update just project description
```

#### Output
```shell
	OK / KO
```
#### 



## al2 projects delete
This Command is used to : 

	- Deletes a project.

#### Command

```shell
al2 projects delete --id=PROJECT_ID
```

#### Examples

```shell
al2 projects delete --id=1 # delete project having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 projects list-stack
This Command is used to : 

	- Gets all stack.

#### Command

```shell
al2 projects list-stack --id=ProjectID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 projects list-stack --id=1 # Gets all stack of project having ID=1
```

#### Output
```shell
	Vertical list
```
#### 


## al2 projects get-stack
This Command is used to : 

	- Gets a stack by their ID.

#### Command

```shell
al2 projects get-stack --id=ProjectID --stackId=StackID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 projects get-stack --id=1 --stackId=2 # Gets stack 2 of project 1
```

#### Output
```shell
	Vertical list
```
#### 


## al2 projects create-stack
This Command is used to : 

	- Creates a stack.

#### Command

* `al2 projects create-stack --id=ProjectID -f FILEPATH ` - Create a stack using file. ( json / yaml)
* `al2 projects create-stack --id=ProjectID ` - Create a stack using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 projects create-stack --id=1 # Create stack with interaction (see create user).
al2 projects create-stack --id=1 -f "C://user/..../file.json" # create stack using json file.
```

#### Output
```shell
	OK / KO
```
#### 

## al2 projects update-stack
This Command is used to : 

	- Updates a stack.

#### Command

* `al2 projects update-stack --id=ProjectID --stackId=STACK_ID -f FILEPATH ` - Update stack using file. ( json / yaml)
* `al2 projects update-stack --id=ProjectID --stackId=STACK_ID ` - Update stack using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to update the resource.

#### Examples

```shell
al2 projects update-stack --id=1 --stackId=2 # Update stack with interaction (see update user).
al2 projects update-stack --id=1 --stackId=2 -f "C://user/..../file.json" # update stack using json file.
```

#### Output
```shell
	OK / KO
```
#### 


## al2 projects delete-stack
This Command is used to : 

	- Deletes a stack.

#### Command

```shell
al2 projects delete-stack --id=PROJECT_ID --stackId=STACK_ID
```

#### Examples

```shell
al2 projects delete-stack --id=1 --stackId=2 # delete stack having ID = 2 from project having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 
