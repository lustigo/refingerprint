import Koa from "koa";
import mount from "koa-mount";
import apiRoute from "./api";
import frontendRoute from "./frontend";

/**
 * Initialize and configure the Router
 */
export default function initRoutes(koa: Koa) {
    koa.use(mount(frontendRoute));
    koa.use(mount("/api", apiRoute));
}
