const gulp = require("gulp");
const rollup = require("rollup");
const resolve = require("@rollup/plugin-node-resolve");
const ts = require("rollup-plugin-typescript2");

gulp.task("default", async function () {
    await bundleContentScript();
    await bundleBackgroundScript();
});

async function bundleBackgroundScript() {
    const bundle = await rollup.rollup({
        input: "./src/io/Background.ts",
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
        file: "./dist/background.js",
        format: "umd",
        name: "bundle",
        sourcemap: false
    });
}

async function bundleContentScript() {
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
}