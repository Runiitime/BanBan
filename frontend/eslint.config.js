import globals from "globals";
import { defineConfig } from "eslint/config";
import { defaultEslintConfigSections } from "./eslint.mjs";

const aliasList = ["app", "data", "infrastructure", "views"];

export default defineConfig([
  ...defaultEslintConfigSections.globalSection,
  {
    files: defaultEslintConfigSections.files,
    ignores: [...defaultEslintConfigSections.ignores, "src/stories/**"],
    extends: [...defaultEslintConfigSections.extends],
    languageOptions: {
      globals: globals.browser,
    },
    plugins: defaultEslintConfigSections.plugins,
    rules: {
      ...defaultEslintConfigSections.rules,

      // Сортировка импортов
      "simple-import-sort/imports": [
        "error",
        {
          groups: [
            // Packages. `react` related packages come first
            ["^react", "^@?\\w"],

            // Alias imports inside repo
            [`^@(${aliasList.join("|")})\\/\\w`],

            // Relative imports, style's imports comes last
            ["^(\\.|\\.\\.)\\/(?!style).*", "^(\\.|\\.\\.)\\/.*style"],
          ],
        },
      ],
    },
  },
]);
