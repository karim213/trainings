# ResourceIntegration

## al2 resources search-vm
This Command is used to : 

	- Finds a resource.

#### Command

```shell
al2 resources search-vm
```

#### Examples

```shell
al2 resources search-vm # Serach virtual machine system Interaction
	Object ID : 1236wf5
	VM instance ID : test
	State : test
	Discovery source : test
```

#### Output
```shell
	Vertical list
```
#### 


## al2 resources list-metrics-vm
This Command is used to : 

	- Gets all resource metrics.

#### Command

```shell
al2 resources list-metrics-vm --systemID=SYSTEM_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 resources list-metrics-vm --systemID=12 # Gets metrics of vm having ID = 1
```

#### Output
```shell
# metric name 1
	Name : metric name 1
	Description : metric description
	Type : metric type
	....

# metric name 2
	Name : metric name 2
	Description : metric description
	Type : metric type
	....

```
#### 


## al2 al2 resources list-tags-vm
This Command is used to : 

	- Gets all resource tags.

#### Command

```shell
al2 resources list-tags-vm --systemID=SYSTEM_ID [--limit=LIMIT] [--sort-by=KEY] [--filter=KEY1,KEY2...]
```

#### List Command Flags
`--limit`=`LIMIT` : Maximum number of resources to list.

`--sort-by`=`KEY` : KEY to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that key. 

`--filter`=`KEY1,KEY2...` : Comma-separated list of resource field key names to display.

#### Examples

```shell
al2 resources list-tags-vm --systemID=123
```

#### Output
```shell
	Table (key, value)
```
#### 

