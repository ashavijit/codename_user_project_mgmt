const express = require('express');
const app = express();
const port = 3000;
const { GraphQLClient, gql } = require('graphql-request');
const URL = 'http://localhost:8080/query';
const graphQLClient = new GraphQLClient(URL, {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json',
    },
});

app.get('/', async (req, res) => {
    res.json({ message: 'Hello World!' });
});

app.get('/id', async (req, res) => {
    const ID = req.query.id;
    const query = gql`    

    query {
        project(input: { id: "${ID}" }) {
            _id
            ownerId
            name
            description
            status
        }
    }
    `;
    const data = await graphQLClient.request(query);
    res.send(data);

});

app.post('/create', async (req, res) => {
    const NAME = req.query.name;
    const mutation = gql`
    mutation{
        createProject(input:
          { ownerId: "649a75623070691b350290c3", name: "${NAME}" , description :"ggg", status: NOT_STARTED}){
          _id
          name
          ownerId
          description
          status
        }
      }
    `;
    const data = await graphQLClient.request(mutation);
});


app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});
