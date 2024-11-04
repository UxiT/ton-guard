require('dotenv').config()
const express = require('express')
const blockchainClient = require('../api/ton/blockchain');

const app = express()
const port = 3000

const blockchainAPI = new blockchainClient();

app.get('/', (req, res) => {
    blockchainAPI.getBlocks()
        .then(blocks => res.json(blocks.data))
})

app.listen(port, () => {
    console.log(`running server on localhost:${port}`)
})