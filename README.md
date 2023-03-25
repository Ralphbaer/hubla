# Sales Service

This repository contains the source code for the Sales Service.

## Architecture

![alt text](./hexagonal-macro.png "Sales Service")

## Requirements

| Name | Version | Notes | Mandatory
|------|---------|---------|---------|
| [golang](https://golang.org/dl/) | >= go1.20.2 | Main programming language | true
| [docker](https://www.docker.com/) | n/a | Used to start local environment providers (MongoDB) | true
| [aws-cli](https://aws.amazon.com/pt/cli/) | v2 | Used to create all AWS Enviroment (Just in case you want to know) | false
| [sh/bash] | depending on OS. Anyway, you should be able do execute any .sh file | Used to lint checks, test processes and some console interface customizations | true
| [make](https://www.gnu.org/software/make/) | depending on OS. Anyway, you should be able do execute make commands to run the project, tests and localenvironment | n/a | true

## Providers

| Name | Version | Notes
|------|---------|---------|
| [aws](https://aws.amazon.com/pt/) | n/a | All the infraestructure are on AWS
| [postgresql](https://www.postgresql.org/) | any stable version | If you want, you can use any postgresql client to access the local database created | true

# Usage
Inside /sales, follow the steps.:

### GitHub Env
```bash
make setup-env            
```

### Start Local Database
- If you want to just test local without any preload data, just run:
```bash
make localenv                     # Start local mongodb empty instance
```

- If you want to preload some test data local, run:
```bash
make localenv-withData            # Start local mongodb with preload data based on .localenv/withData/sales.txt
```

Note: After running any of the above commands, an instance will start on your machine and the thread used in the command line of your choice will be blocked. You need to open a second command line window to run the next command.

### Notes about the Preloaded Dataset
Every time you create a new instance using PostgreSQL preloaded data (make localenv-withData), all the sales data are created again, which means they are not persisted with a preconditioned ID, and PostgreSQL will always generate new IDs. If you want to get a generated ID, it's recommended to open the local instance and choose which ID you want to test (for Get(ID) cases).

### Finally Start Service
After starting the database locally, you are ready to execute the service. Run: 
```bash
make run                  # Start local service on port :3000
```

# Testing

```bash
make test                 # Run all tests
```

## Documentation

Visit [this link](localhost:3000/docs) for API documentation. If you want to access the docs locally, change the host in the URL to localhost:3000. For example: http://localhost:3000/sales/docs

# Deployment

The Sales Service uses a Github Actions pipeline that automatically triggers when code is pushed to the master branch. To facilitate testing, the master branch has no PR approval criteria for the Hubla team. To deploy and test the service, you can make changes to the code or modify the signal.id file inside /sales, commit the changes to master, and wait for Github Actions to deploy the code. 
All necessary provisioning on the AWS Cloud has already been completed and is running. If you encounter any issues accessing the project through the link, feel free to reach out to us, and we can provide more details about AWS provisioning, such as EC2, load balancer, system manager, and secrets, as well as the CircleCI pipeline to help you get started.