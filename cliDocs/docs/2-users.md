# Users

## al2 users list
This Command is used to : 

	- Get all Users.

#### Command

```shell
al2 users list [--all(default : true) | --active(default : false) | --deactivated(default : false) ] [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 users list -all # Gets all users
al2 users list --limit=10 # Gets 10 users
al2 users list --sort-by="name" # Gets all users order by name
al2 users list --sort-by="created_at" --limit=10 # Gets the 10 first users ordred by created date
al2 users list --sort-by="~created_at" --limit=10 # Gets the 10 last users ordred by created date desc
al2 users list --filter="id,profile.name" # Gets IDs, names of all users
al2 users list --filter="id" # Gets IDs of all users
```

#### Output
```shell
	Vertical list
		Id :
		profiles :
			first name :
			last name :
			....
```
#### 
## al2 users get
This Command is used to : 

	- Gets a user.

#### Command

```shell
al2 users get --id=USER_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 users get --id=1 # get user having ID = 1
```

#### Output
```shell
	Vertical list
```
#### 


## al2 users create
This Command is used to : 

	- Creates a user.

#### Command

* `al2 users create -f FILEPATH ` - Create user using file. ( json / yaml)
* `al2 users create --id=USER_ID` - Create user using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 users create -f "C://user/..../file.json" # create team using json file.
al2 users create # Create user using system Interaction.
```

#### Output
```shell
	al2 users create

		First Name : Karim

		Last Name : ABDI

		Email : test@test.com

		Mobile phone : 000003654886

		.....

		User added successfully
```
#### 

## al2 users update
This Command is used to : 

	- Updates a user.

#### Command

* `al2 users update --id=USER_ID -f FILEPATH ` - Update user using file. ( json / yaml)

* `al2 users update --id=USER_ID` - Update user using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to update the resource.


#### Examples

```shell
al2 users update —uid="1" 

		First Name : Karim

		Last Name (ABDI) :                       #<= if we write nothing, the last name will not be updated

		Email (test@test.com): 

		Mobile phone : 000003654886

		.....

		user updated successfully```
```
#### 



## al2 users delete
This Command is used to : 

	- Deletes a user.

#### Command

```shell
al2 users delete --id=USER_ID
```

#### Examples

```shell
al2 users delete --id=1 # delete user having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 users assign-to-folder
This Command is used to : 

	- Creates a user folder relation.

#### Command

```shell
al2 users assign-to-folder --id=USER_ID --folderId=FOLDER_ID
```

#### Examples

```shell
al2 users assign-to-folder --id=1 --folderId=7 # assign user having ID=1 to the folder having ID=7
```

#### Output
```shell
	OK / KO
```
#### 


## al2 users unassign-from-folder
This Command is used to : 

	- Deletes a user folder relation.

#### Command

```shell
al2 users unassign-from-folder --id=USER_ID --folderId=	FOLDER_ID
```

#### Examples

```shell
al2 users  unassign-from-folder --id=1 --folderid=7 # unassign user having ID=1 from the folder having ID=7
```

#### Output
```shell
	OK / KO
```
#### 


## al2 users assigned-folders
This Command is used to : 

	- Gets all folders related to a user.

#### Command

```shell
al2 users assigned-folders --id=USER_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 users assigned-folders --id=1 # List all folders related to the user having ID=1
al2 users assigned-folders --id=1 --filter="id" # List all folders IDs related to the user having ID=1
```

#### Output
```shell
| ID                                   | Name        | Description | parentId                             |
| ------------------------------------ | ----------- | ----------- | ------------------------------------ |
| 22b007c9-c6c6-4100-8e77-6a667c7df856 | Root        | description |                                      |
| f4ef44f4-7dbd-4648-b782-f9df29b7d40c | Test Folder |             | 22b007c9-c6c6-4100-8e77-6a667c7df856 |
```
#### 




## al2 users assign-to-orgunit
This Command is used to : 

	- Creates a user org unit relation.

#### Command

```shell
al2 users assign-to-orgunit --id=USER_ID --orgunitid=ORGUNIT_ID
```

#### Examples

```shell
al2 users assign-to-orgunit --id=1 --orgunitid=8 # assign user having ID=1 to the orgunit having ID=8
```

#### Output
```shell
	OK / KO
```
#### 


## al2 users unassign-from-orgunit
This Command is used to : 

	- Deletes a user org unit relation.

#### Command

```shell
al2 users unassign-from-orgunit --id=USER_ID --orgunitId=ORGUNIT_ID
```

#### Examples

```shell
al2 users unassign-from-orgunit --id=1 --orgunitid=8 # unassign user having ID=1 from the orgunit having ID=8
```

#### Output
```shell
	OK / KO
```
#### 


## al2 users assigned-orgunits
This Command is used to : 

	- Gets all org units related to a user.

#### Command

```shell
al2 users assigned-orgunits --id=USER_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 users assigned-orgunits --id=1 # List all orgunits related to the team having ID=1
al2 users assigned-orgunits --id=1 --filter="id" # List all orgunits IDs related to the team having ID=1
```

#### Output
```shell
| ID                                   | Name       | Description   | folderID                             |
| ------------------------------------ | ---------- | ------------- | ------------------------------------ |
| a20305a0-66f1-407a-a677-a804d2a675e5 | Test Org   | Test org unit | f4ef44f4-7dbd-4648-b782-f9df29b7d40c |
| e42acb67-3a39-47f8-8cce-64d5fa126e96 | Test Org 2 | test org 2    | f4ef44f4-7dbd-4648-b782-f9df29b7d40c |


```
#### 



## al2 users assigned-teams
This Command is used to : 

	- Gets all User Teams.

#### Command

```shell
al2 users assigned-teams --id=USER_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 users assigned-teams --id=1 # List all teams related to the user having ID=1
al2 users assigned-orgunits --id=1 --filter="id" # List all teams IDs related to the user having ID=1
```

#### Output
```shell
	| ID   | Name    | Description |
	| ---- | ------- | ----------- |
	| 1    | UI / UX | Best team   |
	| 2    | DevOps  | DevOps team |
```
#### 

