# serviceAccount

## al2 serviceaccounts list
This Command is used to : 

	- Gets all service accounts.

#### Command

```shell
al2 serviceaccounts list  [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 serviceaccounts list # Gets all service accounts
al2 serviceaccounts list --limit=10 # Gets 10 service accounts
al2 serviceaccounts list --sort-by="name" # Gets all service accounts order by name
al2 serviceaccounts list --sort-by="created_at" --limit=10 # Gets the 10 first service accounts ordred by created date
al2 serviceaccounts list --sort-by="~created_at" --limit=10 # Gets the 10 last service accounts ordred by created date desc
al2 serviceaccounts list --filter="name,parentId" # Gets names, parentId of all service accounts
al2 serviceaccounts list --filter="id" # Gets IDs of all service accounts
```

#### Output
```shell
# Test SA for Catalog
    id: d13f7fd7-f91c-4f2d-926f-0fab5402067e,
    org_id: 00000000-0000-0000-0000-000000000000,
    org_unit_id: 00000000-0000-0000-0000-000000000000,
    name: Test SA for Catalog,
    type: github,
    description: test sa for catalog

# test
    id: ee17f311-6281-4e74-ad6d-e50b94b19ecb,
    org_id: 00000000-0000-0000-0000-000000000000,
    org_unit_id: 00000000-0000-0000-0000-000000000000,
    name: test,
    type: github,
    description: test	

```
#### 


## al2 serviceaccounts get
This Command is used to : 

	- Gets a service account by their ID.

#### Command

```shell
al2 serviceaccounts get --id=SERVICE_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 serviceaccounts get --id="d13f7fd7-f91c-4f2d-926f-0fab5402067e" # get service account having ID = d13f7fd7-f91c-4f2d-926f-0fab5402067e
```

#### Output
```shell
    id: d13f7fd7-f91c-4f2d-926f-0fab5402067e,
    org_id: 00000000-0000-0000-0000-000000000000,
    org_unit_id: 00000000-0000-0000-0000-000000000000,
    name: Test SA for Catalog,
    type: github,
    description: test sa for catalog
```
#### 


## al2 serviceaccounts create
This Command is used to : 

	- Creates a service account.

#### Command

* `al2 serviceaccounts create -f FILEPATH ` - Create service account using file. ( json / yaml)
* `al2 serviceaccounts create` - interaction with the system

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.


#### Examples

```shell
al2 serviceaccounts create -f "C://user/..../file.json" # Creates service account using json file.
al2 serviceaccounts create # Creates service account using system Interaction.
```

#### Output
```shell
	Select service type :
		>Aws
		>Azure
		>Google
		>Github
		>Terraform
		>Vra

	Description : this is description test
	Name: this is nametest
	...
	(show inputs of selected service type)
```
#### 

## al2 serviceaccounts update
This Command is used to : 

	- Updates a service account.

#### Command

* `al2 serviceaccounts update --id=SERVICE_ID -f FILEPATH ` - Update serviceaccounts using file. ( json / yaml)
* `al2 serviceaccounts update --id=SERVICE_ID` - Updates service account using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to update the resource.

#### Examples

```shell

al2 serviceaccounts update --id=1 -f "C://user/..../file.json" # update serviceaccounts using json file.
#interction with system
al2 serviceaccounts update --id=1
	New service Accounts name ("name") :
	New Folder description ("description"): folder test decription
	....
```

#### Output
```shell
	OK / KO
```
#### 



## al2 serviceaccounts delete
This Command is used to : 

	- Deletes a service account.

#### Command

```shell
al2 serviceaccounts delete --id=SERVICE_ID
```

#### Examples

```shell
al2 serviceaccounts delete --id=1 # delete serviceaccounts having ID = 1
```

#### Output
```shell
	OK / KO
```
#### 


## al2 serviceaccounts list-properties
This Command is used to : 

	- Gets a service account properties by their ID.

#### Command

```shell
al2 serviceaccounts list-properties [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 serviceaccounts list-properties --serviceID=1 # Gets all properties of service account
```

#### Output
```shell
	Vertical list
```
#### 



## al2 serviceaccounts list-segments
This Command is used to : 

	- Gets all service account segments.

#### Command

```shell
al2 serviceaccounts list-segments --id=SERVICE_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 serviceaccounts list-segments --id=1 # Gets all segements of service account having ID=1
```

#### Output
```shell
	Vertical list (see list service accounts)
```
#### 


## al2 serviceaccounts get-segment
This Command is used to : 

	- Gets a segment by their ID.

#### Command

```shell
al2 serviceaccounts get-segment --id=SERVICE_ID --segmentID=SEGMENT_ID [--filter=KEY1,KEY2...]
```
#### List Command Flags
`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 serviceaccounts get-segment --id=1 --segmentID=2 # Gets a segment of service accounts by their ID
```

#### Output
```shell
	Vertical list
```
#### 


## al2 serviceaccounts create-segment
This Command is used to : 

	- Creates a service account segment.

#### Command

* `al2 serviceaccounts create-segment --id=SERVICE_ID -f FILEPATH ` - Create segment using file. ( json / yaml)
* `al2 serviceaccounts create-segment --id=SERVICE_ID` - Create segment using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.


#### Examples

```shell
al2 serviceaccounts create-segment --id=1 -f "C://user/..../file.json" # Creates segment using json file.
al2 serviceaccounts create-segment --id=1 # Creates segment using system Interaction.
```

#### Output
```shell
	OK / KO
```
#### 


## al2 serviceaccounts update-segment
This Command is used to : 

	- Updates a service account segment.

#### Command

* `al2 serviceaccounts update-segment --id=SERVICE_ID -f FILEPATH ` - Updates segment using file. ( json / yaml)

* `al2 serviceaccounts update-segment --id=SERVICE_ID --segmentId=SEGMENT_ID` - Updates segment using system Interaction.

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.


## al2 serviceaccounts delete-segment
This Command is used to : 

	- Deletes a service account segment.

#### Command

```shell
al2 serviceaccounts delete-segment --id=SERVICE_ID --segment_id=SEGMENT_ID
```

#### Examples

```shell
al2 serviceaccounts delete-segment --id=1 --segment_id=2 # Deletes segment having ID = 2
```

#### Output
```shell
	OK / KO
```
#### 


## al2 serviceaccounts segments list-regions
This Command is used to : 

	- Gets all region.

#### Command

```shell
al2 serviceaccounts segments list-regions --id=SERVICE_ID --segment_id=SEGMENT_ID  [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.


#### Examples

```shell
al2 serviceaccounts segments list-regions --id=1 --segment_id=2 # Deletes segment having ID = 2
```

#### Output
```shell
	OK / KO
```
####


## al2 serviceaccounts segments create-region
This Command is used to : 

	- Creates a region.

#### Command

* `al2 serviceaccounts segments create-region --id=SERVICE_ID --segmentId=SEGMENT_ID -f FILEPATH ` - Creates region using file. ( json / yaml)
* `al2 serviceaccounts segments create-region --id=SERVICE_ID --segmentId=SEGMENT_ID` - Creates region using system Interaction..

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.

#### Examples

```shell
al2 serviceaccounts segments create-region --id=1 --segmentId=2 -f "C://user/..../file.json" # Creates region using json file.
al2 serviceaccounts segments create-region --id=1 --segmentId=2 # Creates region using system Interaction.
```

#### Output
```shell
	OK / KO
```
#### 


## al2 serviceaccounts segments update-region
This Command is used to : 

	- Updates a region.

#### Command

* `al2 serviceaccounts segments update-region --id=SERVICE_ID --segmentId=SEGMENT_ID --regionID=REGION_ID -f FILEPATH ` - Updates region using file. ( json / yaml)
* `al2 serviceaccounts segments update-region --id=SERVICE_ID --segmentId=SEGMENT_ID --regionID=REGION_ID` - Updates region using system Interaction..

#### List Command Flags
`--fileName`=`FILEPATH` OR `-f FILEPATH` : Filename, or URL to file to use to create the resource.


#### Examples

```shell
al2 serviceaccounts segments update-region --id=1 --segmentId=2 --regionID=1 -f "C://user/..../file.json" # update region using json file.
al2 serviceaccounts segments update-region --id=1 --segmentId=2 --regionID=1 # Update region using system Interaction.
```

#### Output
```shell
	OK / KO
```
#### 


## al2 serviceaccounts segments delete-region
This Command is used to : 

	- Deletes a region.

#### Command

```shell
al2 serviceaccounts segments delete-region --id=SERVICE_ID --segmentId=SEGMENT_ID --regionID=REGION_ID
```

#### Examples

```shell
al2 serviceaccounts segments delete-region --id=1 --segmentId=2 --regionID=1
```

#### Output
```shell
	OK / KO
```
#### 


