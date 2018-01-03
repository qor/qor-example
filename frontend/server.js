const express = require('express');
const next = require('next');
const proxy = require('http-proxy-middleware');

const dev = process.env.NODE_ENV !== 'production';
const app = next({dev});
const handle = app.getRequestHandler();

app
    .prepare()
    .then(() => {
        const server = express();
        server.get('/category/:name', (req, res) => {
            const actualPage = '/category';
            const queryParams = {name: req.params.name};
            app.render(req, res, actualPage, queryParams);
        });

        server.get('/product/:name', (req, res) => {
            const actualPage = '/product_show';
            const queryParams = {name: req.params.name};
            app.render(req, res, actualPage, queryParams);
        });

        server.get('*', (req, res) => {
            return handle(req, res);
        });

        server.listen(3000, err => {
            if (err) throw err;
            console.log('> Ready on http://localhost:3000');
        });
    })
    .catch(ex => {
        console.error(ex.stack);
        process.exit(1);
    });
