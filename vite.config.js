import glob from "glob";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { resolve } from "path";
import { defineConfig } from "vite";

const root = resolve(__dirname, "src");
const outDir = resolve(__dirname, "dist");
const obj = Object.fromEntries(
    glob.sync("src/**/*.html").map((file) => {
        console.log("route: " + file);
        return [
            // This remove `src/` as well as the file extension from each
            // file, so e.g. src/nested/foo.js becomes nested/foo
            path.relative(
                "src",
                file.slice(0, file.length - path.extname(file).length)
            ),
            // This expands the relative paths to absolute paths, so e.g.
            // src/nested/foo becomes /project/src/nested/foo.js
            fileURLToPath(new URL(file, import.meta.url)),
        ];
    })
);

export default defineConfig({
    root,
    base: "./",
    publicDir: resolve(__dirname, "public"),
    build: {
        outDir,
        publicDir: resolve(__dirname, "public"),
        rollupOptions: {
            input: obj,
        },
    },
    resolve: {
        alias: {
            // Here to create a custom path
            "/~style": resolve(__dirname, "./src"),
            "/~script": resolve(__dirname, "./src"),
        },
    },
});
