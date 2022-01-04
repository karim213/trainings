# Catalogs

## al2 catalogs list
his Command is used to : 

	- Get all catalogs.

#### Command

```shell
al2 catalogs list [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 catalogs list -all # Gets all catalogs
al2 catalogs list --limit=10 # Gets 10 catalogs
al2 catalogs list --sort-by="name" # Gets all catalogs order by name
al2 catalogs list --sort-by="created_at" --limit=10 # Gets the 10 first catalogs ordred by created date
al2 catalogs list --sort-by="~created_at" --limit=10 # Gets the 10 last catalogs ordred by created date desc
al2 catalogs list --filter="owner,branch" # Gets branches, names of all catalogs
al2 catalogs list --filter="owner" # Gets owner of all catalogs
```

#### Output
```shell
	Vertical list
```
#### 


## al2 catalogs get
This Command is used to : 

	- Gets a catalog by their ID.

#### Command

```shell
al2 catalogs get --id=CATALOG_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 catalogs get --id=1 # Gets catalogs having ID = 1
```

#### Output
```shell
	Name : UI catalogs name
	Description : catalog description
	Branch : Catalog branch
	....
```
#### 


## al2 catalogs create
This Command is used to : 

	- Creates a catalog.

#### Command

* `al2 catalogs create -f FILEPATH ` - Creates catalogs using file. ( json / yaml)
* `al2 catalogs create` - Creates catalogs using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 teams create -f "C://user/..../file.json" # Creates catalog using json file.
al2 catalogs create # Creates catalogs using system Interaction (see create user)
```

#### Output
```shell
	OK / KO
```
#### 

## al2 catalogs update
This Command is used to : 

	- Updates a catalog.

#### Command

* `al2 catalogs update --id=CATALOG_ID -f FILEPATH ` - Updates catalog using file. ( json / yaml)
* `al2 catalogs update --id=CATALOG_ID` - Updates catalog using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.


#### Examples

```shell
al2 catalogs update --id=1 -f "C://user/..../file.json" # updates team using json file.

#interction with system
al2 catalogs update --id=1
	New catalogs name ("catalog name") :
	New catalogs description ("catalog description"): new catalog description
#in this case we update just catalog description
```

#### Output
```shell
	OK / KO
```
#### 



## al2 catalogs delete
This Command is used to : 

	- Deletes a catalog.

#### Command

```shell
al2 catalogs delete --id=CATALOG_ID
```

#### Examples

```shell
al2 catalogs delete --id=1 # Deletes catalog having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 
