import { defineConfig } from "vite";
import solid from "vite-plugin-solid";

export default defineConfig({
  plugins: [solid()],
  base: '/static/',
  build: {
    outDir: './dist',
    emptyOutDir: true,
    rollupOptions: {
      output: {
        entryFileNames: `index.js`,
        chunkFileNames: `index.js`,
        assetFileNames: `[name].[ext]`,
      },
    },
  },
});
