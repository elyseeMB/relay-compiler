import { defineConfig } from "vite";
import tailwindCSS from "@tailwindcss/vite";
import preact from "@preact/preset-vite";

export default defineConfig({
  plugins: [preact({ babel: { plugins: ["relay"] } }), tailwindCSS()],

  base: "/",
  resolve: {
    alias: {
      react: "preact/compat",
      "react-dom": "preact/compat",
    },
  },
  build: {
    manifest: true,
    outDir: "./pkg/server/public/assets/",
    rollupOptions: {
      input: {
        main: "./assets/main.tsx",
      },
      output: {
        entryFileNames: "[name]-[hash].js",
        assetFileNames: "[name]-[hash][extname]",
        chunkFileNames: "[name]-[hash][extname]",
      },
    },
  },
});
