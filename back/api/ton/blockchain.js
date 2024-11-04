const client = require('../client')

module.exports = class BlockChainApi {
    constructor() {
        this.client = new client(process.env.TON_API_BASE_URL)
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

    getMasterChainBlockShardState(seqno) {
        return this.client.get('v3/masterchainBlockShardState', {
            seqno: seqno,
        })
    }

    getMasterChainBlockShards(seqno, limit, offset) {
        return this.client.get('v3/masterchainBlockShards', {
            seqno: seqno,
            limit: limit,
            offset: offset,
        })
    }

    getMasterChainInfo() {
        return this.client.get('v3/masterchainInfo', {})
    }

    getMessages(params = {}) {
        return this.client.get('v3/messages', params)
    }

    getTransactions(params = {}) {
        return this.client.get('v3/transactions', params)
    }

    getTransactionsByMasterChainBlock(seqno, limit, offset, sort) {
        return this.client.get('v3/transactionsByMasterchainBlock', {
            seqno: seqno,
            limit: limit,
            offset: offset,
            sort: sort
        })
    }

    getTransactionsByMessage(params = {}) {
        return this.client.get('v3/transactionsByMessage', params)
    }
}