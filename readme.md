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