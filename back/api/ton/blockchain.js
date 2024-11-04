const client = require('../client')

module.exports = class BlockChainApi {
    constructor() {
        this.client = new client()
    }

    getAdjacentTransactions(params = {}) {
        this.client.get(
            'v3/adjacentTransactions',
            params
        )
        .then(data => console.log('GET data:', data))
        .catch(error => console.error(error))
    }

    getBlocks(params = {}) {
        return this.client.get('v3/blocks', params)
    }
}