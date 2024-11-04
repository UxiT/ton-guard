const client = require('../client')

module.exports = class AccountsApi {
    constructor() {
        this.client = new client(process.env.TON_API_BASE_URL)
    }

    getAccountStates(params = {}) {
        return this.client.get('v3/accountStates', params)
    }

    getAddressBook(params = {}) {
        return this.client.get('v3/addressBook', params)
    }

    getWalletStates(params = {}) {
        return this.client.get('v3/walletStates', params)
    }
}