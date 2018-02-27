# gosvm

Golang services version manager

# Background

Why development many services, each can used common package, example: logger, DB driver, model structure and etc.

Each sercices using own repository and vendoring (dep, glide or go.mod)

gosvm provide tool for management version through serveral repositoreis

Example:
``` bash
$ pwd
/go/src/github.com/sah4ez/
$ ls -lah
drwxr-xr-x   2 sah4ez  sah4ez    64B 27 feb 21:45 events
drwxr-xr-x   2 sah4ez  sah4ez    64B 27 feb 21:45 logger-wrap
drwxr-xr-x   2 sah4ez  sah4ez    64B 27 feb 21:45 models
drwxr-xr-x   2 sah4ez  sah4ez    64B 27 feb 21:45 mongo
drwxr-xr-x   2 sah4ez  sah4ez    64B 27 feb 21:45 processor
drwxr-xr-x   2 sah4ez  sah4ez    64B 27 feb 21:45 scheduler
```
We have the structure projects:

- events - package implement event for send between sevices or server-client
- logger-warp - wrapper for logger implement own format/logic
- mongo - wrapper for MongoDB driver
- models - structure of models for persist document to MongoDB
- processor - simple service which process input events and store to MongoDB
- schdeudler - execute deferred events

And we can simple describe all structre in `svm.tolm` and management version and update for all services. Examlpe:

```toml
Title = "exmaple-project"
Description = "Example project with library and services"
Version = "1.0.0"
BasePath = "github.com/sah4ez"

[[SubProject]]
Title = "events"
Description = "package implement event for send between sevices or server-client"
Version = "1.1.0"

[[SubProject]]
Title = "logger-warp"
Description = "wrapper for logger implement own format/logic"
Version = "0.2.1"

[[SubProject]]
Title = "mongo"
Description = "wrapper for MongoDB driver"
Version = "0.2.1"

[[SubProject]]
Title = "models"
Description = "structure of models for persist document to MongoDB"
Version = "0.2.1"

[[SubProject]]
Title = "processor"
Description = "simple service which process input events and store to MongoDB"
Version = "0.2.1"

[[SubProject]]
Title = "schdeudler"
Description = "execute deferred events"
Version = "0.2.1"
```

# Commands

| name | description |
|-----:|-------------|
| list | formatted output svm.toml |
| libs | analyse of usage packages in all packages and formatted output |

Example commands:
```bash
$ ls
svm.toml

$ gosvm list
Title:		 exmaple-project
Description:	 Example project with library and services
Version:	 1.0.0
	
	SubPackages:
	
	events@1.1.0
	logger-warp@0.2.1
	mongo@0.2.1
	models@0.2.1
	processor@0.2.1
	schdeudler@0.2.1

$ gosvm libs
Title:		 exmaple-project
Description:	 Example project with library and services
Version:	 1.0.0

====================Libs====================

	 github.com/sah4ez/events
	 !!!2 differnt vesrion are used!!!
		^1.2.0 :
			processor
		^1.2.2 :
			scheduler

	 github.com/sah4ez/logger-wrap
	 !!!4 differnt vesrion are used!!!
		^1.0.0 :
			events
		^1.3.0 :
			models
		^3.2.0 :
			processor
		^4.2.1 :
			scheduler

	 github.com/sah4ez/models
	 !!!2 differnt vesrion are used!!!
		^1.0.0 :
			processor
		^1.2.1 :
			scheduler

	 github.com/sah4ez/mongo
	 !!!3 differnt vesrion are used!!!
		^2.3.1 :
			scheduler
		^2.3.0 :
			models
		^2.2.0 :
			processor

	 github.com/sirupsen/logrus
		master :
			logger-wrap
```


# dependecy manager

Supported:

- dep
- glide
- ...
