const express = require('express');
const app = express();
const port = 3000;
const { GraphQLClient, gql } = require('graphql-request');
const URL = 'http://localhost:8080/graphql';
const graphQLClient = new GraphQLClient(URL, {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json',
    },
});
const query = gql`
  query {
    project(input: { id: "649a75623070691b350290c3" }) {
      _id
      ownerId
      name
      description
      status
    }
  }
`;
JSON.stringify(query);

app.get('/', async (req, res) => {
    try {
        const data = await graphQLClient.request(query);
        res.send(data);
    } catch (error) {
        console.error('GraphQL Error:', error.response.errors);
        res.status(500).send('Internal Server Error');
    }
});


app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});
