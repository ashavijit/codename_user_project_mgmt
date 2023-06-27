CreatOwner
```graphql
mutation{
  createOwner(input:
    { name: "avijit" , email: "avijit@gamif.vo", phone : "8250325238"}){
    _id
    name
    email
    phone
  }
}
```
createProject

```graphql
mutation{
  createProject(input:
    { ownerId: "649a75623070691b350290c3", name: "okakkaka" , description :"ojcmfmvldsvmds", status: NOT_STARTED}){
    _id
    name
    ownerId
    description
    status
  }
}
```
```graphql
query {
  project(input: { id: "PROJECT_ID_HERE" }) {
    _id
    ownerId
    name
    description
    status
  }
}
```

curl -X POST -H "Content-Type: application/json" -d '{"query": "{ project(input: { id: \"PROJECT_ID_HERE\" }) { _id ownerId name description status } }"}' http://localhost:8080/


curl -X POST -H "Content-Type: application/json" -d '{ 
  "query": "mutation { createOwner(input: { name: \"avijit\", email: \"avijit@gamif.vo\", phone: \"8250325238\" }) { _id name email phone } }"
}' https://codename-user.onrender.com/
