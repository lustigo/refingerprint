// Load the configuration
require("dotenv").config();

import koa from "koa";
import https from "https";
import fs from "fs";
import initRoutes from "./router";

const app = new koa();

/*
 * Add the Routes
 */
initRoutes(app);

/*
 * Start the server
 */
const options = {
    key: fs.readFileSync(process.env.TLS_KEY_PATH as string),
    cert: fs.readFileSync(process.env.TLS_CERT_PATH as string)
};
https.createServer(options, app.callback()).listen(process.env.PORT);
console.log(`Listening on ${process.env.PORT}`);
