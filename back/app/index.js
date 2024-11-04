require('dotenv').config()
const http = require('http');
const blockchainClient = require('../api/ton/blockchain');

const host = '127.0.0.1';
const port = 7000;

function notFound(res) {
    res.statusCode = 404;
    res.setHeader('Content-Type', 'text/plain');
    res.end('Not found\n');
}

const blockchainAPI = new blockchainClient();

const server = http.createServer((req, res) => {
    switch (req.method) {
        case 'GET': {
            switch (req.url) {
                case '/': {
                    res.statusCode = 200;
                    res.setHeader(
                        'Content-Type',
                        'text/plain'
                    );

                    blockchainAPI.getBlocks()
                        .then(
                            (blocks) => res.end(JSON.stringify(blocks.data))
                        );
                    break;
                }
                default: {
                    break;
                }
            }
            break;
        }
        default: {
            notFound(res);
            break;
        }
    }
});

server.listen(port, host, () => {
    console.log(`Server listens http://${host}:${port}`);
});