# Folders

## al2 folders list
This Command is used to : 

	- Gets all folders.

#### Command

```shell
al2 folders list [--tree] [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

`--tree`=`LIMIT` : Display resources as tree.

#### Examples

```shell
al2 folders list # Gets all folders
al2 folders list --limit=10 # Gets 10 folders
al2 folders list --sort-by="name" # Gets all folders order by name
al2 folders list --sort-by="created_at" --limit=10 # Gets the 10 first folders ordred by created date
al2 folders list --sort-by="~created_at" --limit=10 # Gets the 10 last folders ordred by created date desc
al2 folders list --filter="name,parentId" # Gets names, parentId of all folders
al2 folders list --filter="id" # Gets IDs of all folders
al2 folders list --tree # Gets all folders As tree
```

#### Output
```shell
| ID                                   | Name        | Description | ParentID                             |
| ------------------------------------ | ----------- | ----------- | ------------------------------------ |
| 22b007c9-c6c6-4100-8e77-6a667c7df856 | Root        |             |                                      |
| f4ef44f4-7dbd-4648-b782-f9df29b7d40c | Test Folder |             | 22b007c9-c6c6-4100-8e77-6a667c7df856 |

```
#### 


## al2 folders get
This Command is used to : 

	- Get a folder by their ID.

#### Command

```shell
al2 folders get --id=FOLDER_ID  [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 folders get --id=1 # Gets folder having ID = 1
```

#### Output
```shell
	Name : Test Folder
	Description : description folder
	ParentID : 22b007c9-c6c6-4100-8e77-6a667c7df856
```
#### 


## al2 folders create
This Command is used to : 

	- Creates a folder.

#### Command

* `al2 folders create -f FILEPATH ` - Creates folder using file. ( json / yaml)
* `al2 folders create --name=NAME --description=DESCRIPTION --parentId=ORG_ID ` - Creates folder using flags.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 folders create --name="folder 3" --description="description 3" --parentId="22b007c9-c6c6-4100-8e77-6a667c7df856" # Creates folder using flags.
al2 folders create -f "C://user/..../file.json" # creates folder using json file.
al2 folders create # Create folder using system Interaction.
```

#### Output
```shell
	OK / KO
```
#### 

## al2 folders update
This Command is used to : 

	- Updates a folder.

#### Command

* `al2 folders update --id=FOLDER_ID -f FILEPATH ` - Updates folder using file. ( json / yaml)

* `al2 folders update --id=FOLDER_ID --name=NAME --description=DESCRIPTION --parentID=Parent_ID` - Updates folder using flags.

* `al2 folders update --id=FOLDER_ID` - Updates folder using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 folders update --id=1 --name "folder test" --description "folder test decription" # updates name and description of the folder having id = 1
al2 folders update --id=1 --name "folder test" # updates name of the folder having id = 1
al2 folders update --id=1 --description "folder test decription" # updates description of the folder having id = 1
al2 folders update --id=1 -f "C://user/..../file.json" # updates folder using json file.

#interction with system
al2 folders update --id=1
	New Folder name ("folder 3") :
	New Folder description ("description 3"): folder test decription
	New Folder ParentID ("22b007c9-c6c6-4100-8e77-6a667c7df856") :
#in this case we update just folder description
```

#### Output
```shell
	OK / KO
```
#### 



## al2 folders delete
This Command is used to : 

	- Deletes a folder.

#### Command

```shell
al2 folders delete --id=FOLDER_ID
```

#### Examples

```shell
al2 folders delete --id=1 # deletes folder having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 folders childrens
This Command is used to : 

	- Gets a folder subtree by their ID (excluding the folder itself).

#### Command

```shell
al2 folders childrens --id=FolderID [--tree] [--depth=VALUE] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### List Command Flags
`--tree`=`LIMIT` : Display resources as tree.
`--depth`=`VALUE` : Depth of the subtree to return (omit to return the entire subtree).


#### Examples

```shell
al2 folders childrens --id=1 # Gets all childrens of folder having ID=1
```

#### Output
```shell
	Table OR tree
```
#### 


## al2 folders list-orgunits
This Command is used to : 

	- Gets Gets all org units.

#### Command

```shell
al2 folders list-orgunits --id=FolderID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 folders list-orgunits --id=1 # Gets all orgunits of folder having ID=1
```

#### Output
```shell
| ID                                   | Name        | Description | FolderID                             |
| ------------------------------------ | ----------- | ----------- | ------------------------------------ |
| a20305a0-66f1-407a-a677-a804d2a675e5 | Test Org    |Test org unit| f4ef44f4-7dbd-4648-b782-f9df29b7d40c |
| f4ef44f4-7dbd-4648-b782-f9df29b7d40c | Test Folder |test org 2   | f4ef44f4-7dbd-4648-b782-f9df29b7d40c |

```
#### 


## al2 folders create-orgunit
This Command is used to : 

	- Creates a org unit.

#### Command

```shell
al2 folders create-orgunit
```

#### Examples

```shell
al2 folders create-orgunit # Creates orgunit using system Interaction.
```

#### Output
```shell
	OK / KO (need more information)
```
#### 