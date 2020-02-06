import Koa from "koa";
import serve from "koa-static";

/**
 * Includes the Frontend
 */
const app = new Koa();

app.use(serve(process.env.APP_PATH as string));

export default app;