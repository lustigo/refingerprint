import koa from "koa";
const app = new koa();

app.use(require("koa-static")("../frontend/dist/"));

app.listen(8080);