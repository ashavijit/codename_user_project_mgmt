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


app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});
