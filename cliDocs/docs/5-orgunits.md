# Orgunits

## al2 orgunits list
This Command is used to : 

	- Gets every org unit in all folders.

#### Command

```shell
al2 orgunits list [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 orgunits list # Gets all orgunits
al2 orgunits list --limit=10 # Gets 10 orgunits
al2 orgunits list --sort-by="name" # Gets all orgunits order by name
al2 orgunits list --sort-by="created_at" --limit=10 # Gets the 10 first orgunits ordred by created date
al2 orgunits list --sort-by="~created_at" --limit=10 # Gets the 10 last orgunits ordred by created date desc
al2 orgunits list --filter="name,folderId" # Gets names, folderId of all orgunits
al2 orgunits list --filter="id" # Gets IDs of all orgunits
```

#### Output
```shell
| ID                                   | Name        | Description | folderID                             |
| ------------------------------------ | ----------- | ----------- | ------------------------------------ |
| a20305a0-66f1-407a-a677-a804d2a675e5 | Test Org    |Test org unit| f4ef44f4-7dbd-4648-b782-f9df29b7d40c |
| e42acb67-3a39-47f8-8cce-64d5fa126e96 | Test Org 2  |test org 2   | f4ef44f4-7dbd-4648-b782-f9df29b7d40c |

```
#### 


## al2 orgunits get
This Command is used to : 

	- Gets a org unit by their ID.

#### Command

```shell
al2 orgunits get --id=ORG_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 orgunits get --id=1 # Gets orgunit having ID = 1
```

#### Output
```shell
	Name : Test Org
	Description : Test org unit
	FolderID : f4ef44f4-7dbd-4648-b782-f9df29b7d40c
```
#### 


## al2 orgunits update
This Command is used to : 

	- Updates a org unit.

#### Command

* `al2 orgunits update --id=ORG_ID -f FILEPATH ` - Updates orgunit using file. ( json / yaml)

* `al2 orgunits update --id=ORG_ID --name=NAME --description=DESCRIPTION --folderID=FOLDER_ID` - Updates orgunit using flags.

* `al2 orgunits update --id=ORG_ID` - Updates orgunit using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to update the resource.


#### Examples

```shell
al2 orgunits update --id=1 --name "orgunit test" --description "orgunit test decription" # updates name and description of the orgunit having id = 1
al2 orgunits update --id=1 --name "orgunit test" # updates name of the orgunit having id = 1
al2 orgunits update --id=1 --description "orgunit test decription" # updates description of the orgunit having id = 1
al2 orgunits update --id=1 -f "C://user/..../file.json" # updates orgunit using json file.

#interction with system
al2 orgunits update --id=1
	New Orgunit name ("orgunit test") :
	New Orgunit description ("description test"): orgunit test decription
	New Orgunit folderID ("1") :
#in this case we update just orgunit description
```

#### Output
```shell
	OK / KO
```
#### 



## al2 orgunits delete
This Command is used to : 

	- Deletes a org unit.

#### Command

```shell
al2 orgunits delete --id=ORG_ID
```

#### Examples

```shell
al2 orgunits delete --id=1 # Deletes orgunit having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 orgunits assigned-folder
This Command is used to : 

	- Gets org unit folder by their ID .

#### Command

```shell
al2 orgunits assigned-folder --id=ORG_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 orgunits assigned-folder --id=1 # Gets folder that containt orgunit having ID=1
```

#### Output
```shell
	Vertical list
```
#### 


## al2 orgunits list-projects
This Command is used to : 

	- Gets org unit projects .

#### Command

```shell
al2 orgunits list-projects --id=ORG_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 orgunits list-projects --id=1 # Gets all project assigned to the orgunit having ID=1
```

#### Output
```shell
| ID                                   | Name        | Description |
| ------------------------------------ | ----------- | ----------- |
| 22b007c9-c6c6-4100-8e77-6a667c7df856 | project1    | description |
| f4ef44f4-7dbd-4648-b782-f9df29b7d40c | project2    | description |

```
#### 


## al2 orgunits assign-to-project
This Command is used to : 

	- Creates a org unit project.

#### Command

```shell
al2 orgunits assign-to-project --id=ORG_ID --projectId=PROJECT_ID
```

#### Examples

```shell
al2 orgunits assign-to-project --id=1 --projectId=2 # assign orgunit having ID=1 to the project having ID=2
```

#### Output
```shell
	OK / KO
```
#### 


## al2 orgunits unassign-from-project
This Command is used to : 

	- Deletes a org unit project.

#### Command

```shell
al2 orgunits unassign-from-project --id=ORG_ID --projectId=PROJECT_ID
```

#### Examples

```shell
al2 orgunits unassign-from-project --id=1 --projectId=2 # unassign an orgunit having ID=1 from the project having ID=2
```

#### Output
```shell
	OK / KO
```
#### 




## al2 orgunits list-serviceaccounts
This Command is used to : 

	- Gets all org unit service accounts.

#### Command

```shell
al2 orgunits list-serviceaccounts --id=ORG_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 orgunits list-serviceaccounts --id=1 # Gets all service accounts assigned to the orgunit having ID=1
```

#### Output
```shell
	Vertical List (change the api to return all information about service account)

```
#### 


## al2 orgunits assign-to-serviceaccount
This Command is used to : 

	- Registers a service account with an org unit.

#### Command

```shell
al2 orgunits assign-to-serviceaccount --id=ORG_ID --serviceaccountId=SERVICE_ID
```

#### Examples

```shell
al2 orgunits assign-to-serviceaccount --id=1 --serviceaccountId=2 # assign orgunit having ID=1 to the service account having ID=2
```

#### Output
```shell
	OK / KO
```
#### 


## al2 orgunits unassign-from-serviceaccount
This Command is used to : 

	- Deletes a org unit service account reference.

#### Command

```shell
al2 orgunits unassign-from-serviceaccount --id=ORG_ID --serviceaccountId=SERVICE_ID
```

#### Examples

```shell
al2 orgunits unassign-from-serviceaccount --id=1 --serviceaccountId=2 # unassign an orgunit having ID=1 from the service account having ID=2
```

#### Output
```shell
	OK / KO
```
#### 