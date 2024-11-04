const axios = require('axios');

module.exports = class BaseClient {
  constructor() {
    const baseURL = process.env.TON_API_BASE_URL
    const timeout = 5000

    console.log(baseURL);

    this.client = axios.create({
      baseURL,
      timeout,
      headers: {
            'Content-Type': 'application/json',
            'X-Api-Key': process.env.TONWEB_API_KEY
        },
    });
  }

  // GET request
  async get(endpoint, params = {}) {
    try {
      return await this.client.get(endpoint, { params });
    } catch (error) {
      this.handleError(error);
    }
  }

  // POST request
  async post(endpoint, data = {}) {
    try {
      const response = await this.client.post(endpoint, data);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  // Private method to handle errors
  handleError(error) {
    if (error.response) {
      // Server responded with a status other than 2xx
      console.error(`API Error: ${error.response.status} - ${error.response.data}`);
    } else if (error.request) {
      // Request was made, but no response was received
      console.error('No response received:', error.request);
    } else {
      // Something else caused an error
      console.error('Error:', error.message);
    }
    throw error; // Re-throw error after logging
  }
}
