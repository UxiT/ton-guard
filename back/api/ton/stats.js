const client = require('../client')

module.exports = class StatsApi {
    constructor() {
        this.client = new client(process.env.TON_API_BASE_URL)
    }

    getTopAccountsByBalance(params = {}) {
        return this.client.get('v3/topAccountsByBalance', params)
    }
}