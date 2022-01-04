# Teams

## al2 teams list
This Command is used to : 

	- Get all teams.

#### Command
```shell
al2 teams list  [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 teams list # Gets all teams
al2 teams list --limit=10 # Gets 10 teams
al2 teams list --sort-by="name" # Gets all teams order by name
al2 teams list --sort-by="created_at" --limit=10 # Gets the 10 first teams ordred by created date
al2 teams list --sort-by="~created_at" --limit=10 # Gets the 10 last teams ordred by created date desc
al2 teams list --filter="id,name" # Gets IDs, names of all teams
al2 teams list --filter="id" # Gets IDs of all teams
```

#### Output

| ID   | Name    | Description |
| ---- | ------- | ----------- |
| 1    | UI / UX | Best team   |
| 2    | DevOps  | DevOps team |


## al2 teams get
This Command is used to : 

	- Gets a Team by ID.

#### Command

```shell
al2 teams get --ID=TEAM_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 teams get --ID=1 # get team having ID = 1
```

#### Output
```shell
	Name : UI / UX
	Description : Best team
```
#### 


## al2 teams create
This Command is used to : 

	- Creates a Team.

#### Command

* `al2 teams create -f FILEPATH ` - Create team using file. ( json / yaml)
* `al2 teams create --name=NAME --description=DESCRIPTION ` - Create team using flags.
* `al2 teams create` - Create team using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 teams create --name "UI / UX" --description "Best team" # Create team using flags.
al2 teams create -f "C://user/..../file.json" # create team using json file.
```

#### Output
```shell
	OK / KO
```
#### 

## al2 teams update
This Command is used to : 

	- Updates a Team.


#### Command

* `al2 teams update --ID=TEAM_ID -f FILEPATH ` - Update team using file. ( json / yaml)

* `al2 teams update --ID=TEAM_ID --name=NAME --description=DESCRIPTION` - Update team using flags.

* `al2 teams update --ID=TEAM_ID` - Update team using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to update the resource.


#### Examples

```shell
al2 teams update --ID=1 --name "UI / UX 2" --description "Best team ++" # update name and description of the team having id = 1
al2 teams update --ID=1 --name "UI / UX 2" # update name of the team having id = 1
al2 teams update --ID=1 --description "Best team ++" # update description of the team having id = 1
al2 teams update --ID=1 -f "C://user/..../file.json" # update team using json file.

#interction with system
al2 teams update --ID=1
	New team name ("UI / UX") :
	New team description ("Best team"): Best team ++ 
#in this case we update just team description
```

#### Output
```shell
	OK / KO
```
#### 



## al2 teams delete
This Command is used to : 

	- delete team by ID.

#### Command

```shell
al2 teams delete --ID=TEAM_ID
```

#### Examples

```shell
al2 teams delete --ID=1 # delete team having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 teams assign-to-folder
This Command is used to : 

	- Creates a team folder relation.

#### Command

```shell
al2 teams assign-to-folder --ID=TEAM_ID --folderId=FOLDER_ID 
```

#### Examples

```shell
al2 teams assign-to-folder --id=1 --folderId=23 # assign team having ID=1 to the folder having ID=23
```

#### Output
```shell
	OK / KO
```
#### 


## al2 teams unassign-from-folder
This Command is used to : 

	- Deletes a team folder relation.

#### Command

```shell
al2 teams unassign-from-folder --id=TEAM_ID --folderId=	FOLDER_ID
```

#### Examples

```shell
al2 teams unassign-from-folder --id=1 --folderd=23 # unassign team having ID=1 from the folder having ID=23
```

#### Output
```shell
	OK / KO
```
#### 


## al2 teams assigned-folders
This Command is used to : 

	- Gets all folders related to a team.

#### Command

```shell
al2 teams assigned-folders --id=TEAM_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 teams assigned-folders --id=1 # List all folders related to the team having ID=1
al2 teams assigned-folders --id=1 --filter="id" # List all folders IDs related to the team having ID=1
```

#### Output
```shell
| ID                                   | Name        | Description | parentId                             |
| ------------------------------------ | ----------- | ----------- | ------------------------------------ |
| 22b007c9-c6c6-4100-8e77-6a667c7df856 | Root        | description |                                      |
| f4ef44f4-7dbd-4648-b782-f9df29b7d40c | Test Folder |             | 22b007c9-c6c6-4100-8e77-6a667c7df856 |
```
#### 




## al2 teams assign-to-orgunit
This Command is used to : 

	- Creates a team org unit relation.

#### Command

```shell
al2 teams assign-to-orgunit --id=TEAM_ID --orgunitid=ORG_ID
```

#### Examples

```shell
al2 teams assign-to-orgunit --id=1 --orgunitid=12 # assign team having ID=1 to the orgunit having ID=12
```

#### Output
```shell
	OK / KO
```
#### 


## al2 teams unassign-from-orgunit
This Command is used to : 

	- Deletes a team org unit relation.

#### Command

```shell
al2 teams unassign-from-orgunit --id=TEAM_ID --orgunitid=ORGUNIT_ID
```

#### Examples

```shell
al2 teams unassign-from-orgunit --id=1 --orgunitid=12 # unassign team having ID=1 from the orgunit having ID=12
```

#### Output
```shell
	OK / KO
```
#### 


## al2 teams assigned-orgunits
This Command is used to : 

	- Gets all org units related to a team.

#### Command

```shell
al2 teams assigned-orgunits --id=TEAM_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 teams assigned-orgunits --id=1 # List all orgunits related to the team having ID=1
al2 teams assigned-orgunits --id=1 --filter="id" # List all orgunits IDs related to the team having ID=1
```

#### Output
```shell
| ID                                   | Name       | Description   | folderID                             |
| ------------------------------------ | ---------- | ------------- | ------------------------------------ |
| a20305a0-66f1-407a-a677-a804d2a675e5 | Test Org   | Test org unit | f4ef44f4-7dbd-4648-b782-f9df29b7d40c |
| e42acb67-3a39-47f8-8cce-64d5fa126e96 | Test Org 2 | test org 2    | f4ef44f4-7dbd-4648-b782-f9df29b7d40c |


```
#### 



## al2 teams add-user
This Command is used to : 

	- Creates a Team User relation.

#### Command

```shell
al2 teams add-user --id=TEAM_ID --userID=USER_ID
```

#### Examples

```shell
al2 teams add-user --id=1 --userID=2 # assign user having ID=2 to the team having ID=1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 teams remove-user
This Command is used to : 

	- Deletes a Team User relation.

#### Command

```shell
al2 teams remove-user --id=TEAM_ID --userId=USER_ID
```

#### Examples

```shell
al2 teams remove-user --id=1 --userId=2 # unassign user having ID=2 from the team having ID=1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 teams assigned-users
This Command is used to : 

	- Gets a Team users by their ID. 

#### Command

```shell
al2 teams assigned-users --id=TEAM_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 teams assigned-users --id=1 # List all users related to the team having ID=1
al2 teams assigned-users --id=1 --filter="id,profile.firstName" # List all user IDs and first names related to the team having ID=1
```

#### Output
```shell
Vertical LIst


```
#### 