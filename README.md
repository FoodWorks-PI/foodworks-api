# Foodworks Backend
## Getting Started
### Prerequisites
 - Go + PATH environment variable
 - Docker + docker-compose (if using docker for mac or docker for windows you should already have this)
 - Clone of [FoodWorks Auth](https://github.com/FoodWorks-PI/foodworks-auth)
### Running with docker
**TODO**
### Running locally 
- Start the [FoodWorks Auth](https://github.com/FoodWorks-PI/foodworks-auth) services following the instructions on the repo.
- VsCode:
    -  Run -> Start/Run without debugging
- Goland: 
    - Install the Envfile plugin 
    - Open main.go and click on the green arrow, afterwards you should be able to run the project from the toolbar.
- Console:
    - **TODO**
- If you see a message relating to db connection, make sure you are running the auth services and loading `.env.development` when running this project (this should be handle by VsCode or Goland).
## Contributing
### Project Structrure
```
:.
├───.github              # CI/CD for github actions
├───docker               # Docker build files
├───internal    
│   ├───api              # Http and route-controller mapping
│   ├───auth
│   ├───generated        # Genrated code for Ent and GQLGen, do not modify 
│   │   ├───ent
│   │   └───graphql
│   ├───platform         # Support code 
│   └───resolver         # GraphQL
├───schema                  
│   ├───db               # DB schema consumed by Ent
│   └───graphql          # GraphQL schema
└───test                 # Integration tests
```
## Making changes
 1. Create a branch for your contribution `/username/feature` or `/username/feature/CLICKUP_ID`
 1. Add your code
 1. If schema was modified, run `go generate ./...` from the project root. This will re-generate both Ent and GQLGen
 1. Optional: Add an integration test for the feature
 1. Test your code using `go test./...`
 1. Push your changes and submit a PR
### Pre-commit
For format standarization, this repo uses [Pre-Commit](https://pre-commit.com/#installation), if your code does not pass the format checks, the PR will be automatically rejected. 
You can run the checks locally by installing the tool: [Pre-Commit Installation](https://pre-commit.com/#installation) and running `pre-commit` on the root of the project. After making sure your project passes the tests, you can commit, push your changes, and submit a PR.