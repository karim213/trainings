# Mappings

## al2 mappings list
This Command is used to : 

	- Gets all mappings.

#### Command

```shell
al2 mappings list --orgunitID=ORGUNIT_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 mappings list --orgunitID=1 -all # Gets all mappings of orgunit having ID=1
al2 mappings list --orgunitID=1 --limit=10 # Gets 10 mappings of orgunit having ID=1
al2 mappings list --orgunitID=1 --sort-by="name" # Gets all mappings order by name of orgunit having ID=1
al2 mappings list --orgunitID=1 --sort-by="created_at" --limit=10 # Gets the 10 first mappings ordred by created date of orgunit having ID=1
al2 mappings list --orgunitID=1 --sort-by="~created_at" --limit=10 # Gets the 10 last mappings ordred by created date desc of orgunit having ID=1
al2 mappings list --orgunitID=1 --filter="name,stackResource" # Gets names and stackResources of all mappings of orgunit having ID=1
al2 mappings list --orgunitID=1 --filter="orgUnitId" # Gets orgUnitId of all mappings of orgunit having ID=1
```

#### Output
```shell
	Vertical list
```
#### 


## al2 mappings get
This Command is used to : 

	- Gets a mapping by their ID.

#### Command

```shell
al2 mappings get --id=MAPPING_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 mappings get --id=1 # Gets mapping having ID = 1
```

#### Output
```shell
	Name : mapping name
	OrgUnitID : orgunit ID
	StackResource : stack resource
	....
```
#### 


## al2 mappings create

#### Command
This Command is used to : 

	- Creates a mapping.

* `al2 mappings create -f FILEPATH ` - Creates mapping using file. ( json / yaml)
* `al2 mappings create` - Creates mapping using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 mappings create -f "C://user/..../file.json" # create mapping using json file.
al2 mappings create # Create mapping using system Interaction (see create user)
```

#### Output
```shell
	OK / KO
```
#### 

## al2 catalogs update
This Command is used to : 

	- Updates a mapping.

#### Command

* `al2 mappings update --id=MAPPING_ID -f FILEPATH ` - Update mapping using file. ( json / yaml)
* `al2 mappings update --id=MAPPING_ID` - Update mapping using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to update the resource.

#### Examples

```shell
al2 mappings update --id=1 -f "C://user/..../file.json" # update mapping using json file.

#interction with system
al2 mappings update --id=1
	New mapping name ("mapping name") :
	New mapping StackResource ("Stack resource"): new stack resource
#in this case we update just mapping StackResource
```

#### Output
```shell
	OK / KO
```
#### 



## al2 mappings delete
This Command is used to : 

	- delete a mapping.

#### Command

```shell
al2 mappings delete --id=MAPPING_ID
```

#### Examples

```shell
al2 mappings delete --id=1 # delete mapping having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 
