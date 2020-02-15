const gulp = require("gulp");
const rollup = require("rollup");
const resolve = require("@rollup/plugin-node-resolve");
const ts = require("rollup-plugin-typescript2");

gulp.task("default", async function () {
    const bundle = await rollup.rollup({
        input: "./src/index.ts",
        plugins: [
            resolve({
                extensions: [".ts"]
            }),
            ts({
                noResolve: false,
                moduleResolution: "node",
                include: ["src/**/*.ts", "src/*.ts"]
            })
        ],
    });

    await bundle.write({
        file: "./dist/bundle.js",
        format: "umd",
        name: "bundle",
        sourcemap: false
    });
});