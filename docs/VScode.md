# Setting.json

```json
{
  "editor.fontSize": 16,
  "editor.formatOnPaste": true,
  "editor.fontLigatures": false,
  "editor.formatOnSave": true,
  "editor.formatOnType": true,
  "editor.links": false,
  "editor.fontWeight": "normal",
  "editor.codeActionsOnSave": {
    "source.organizeImports": true
    // "source.fixAll": true  // 注释掉不然保存很卡
  },
  "editor.quickSuggestionsDelay": 0,
  "editor.rename.enablePreview": false,
  "editor.unicodeHighlight.nonBasicASCII": false,
  "remote.SSH.showLoginTerminal": true,
  "remote.downloadExtensionsLocally": true,
  "remote.SSH.defaultForwardedPorts": [],
  "remote.SSH.remotePlatform": {
    "ip": "linux"
  },
  "remote.SSH.configFile": "C:\\Users\\admin\\.ssh\\config",
  "workbench.colorCustomizations": {
    //覆盖当前所选颜色主题的颜色
    "editor.background": "#241d1d", //编辑器背景色
    "editor.selectionBackground": "#6b5244", //用户选中代码段的颜色
    "editor.findMatchBackground": "#ff0000", //当前搜索匹配的颜色
    "editor.findMatchHighlightBackground": "#ff00ff", //其他搜索匹配项的颜色
    "editor.findRangeHighlightBackground": "#ff9900", //限制搜索范围的颜色
    "editor.lineHighlightBackground": "#48314e", //光标所在行高亮内容的背景颜色
    "editor.lineHighlightBorder": "#704b36" //光标所在行四周边框的背景颜色
  },
  "workbench.preferredDarkColorTheme": "Monokai",
  "workbench.preferredHighContrastColorTheme": "Abyss",
  "workbench.preferredLightColorTheme": "Monokai",
  "workbench.editor.enablePreviewFromQuickOpen": true,
  "workbench.startupEditor": "none",
  "window.openFoldersInNewWindow": "on",
  "window.newWindowDimensions": "offset",
  "security.workspace.trust.untrustedFiles": "open",
  "[jsonc]": {
    "editor.quickSuggestions": {
      "strings": true
    },
    "editor.suggest.insertMode": "replace"
  },
  "json.schemaDownload.enable": false,
  "files.associations": {
    "*.json": "jsonc"
  },
  "files.autoSave": "onFocusChange",
  "explorer.confirmDelete": false,
  "search.followSymlinks": false,
  "[go]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "go.useLanguageServer": true,
  "go.inferGopath": false,
  "[go.mod]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "python.languageServer": "Default",
  "python.linting.enabled": true,
  "python.formatting.provider": "yapf",
  "[python]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.defaultFormatter": "ms-python.python"
  },
  "python.formatting.yapfArgs": [
    "--style={column_limit=1024}"
    // "max-line-length:120", // 解决python代码自动换行问题
  ],
  "python.linting.pylintEnabled": false,
  "git.openRepositoryInParentFolders": "always",
  "editor.unicodeHighlight.allowedLocales": {
    "zh-hant": true
  },
  "extensions.ignoreRecommendations": true,
  "C_Cpp.intelliSenseEngineFallback": "disabled",
  "C_Cpp.intelliSenseEngine": "Tag Parser",
  "workbench.editor.enablePreview": false,
  "explorer.openEditors.minVisible": 1,
  "window.zoomLevel": 1,
  "editor.tokenColorCustomizations": {
    // "comments": "#FFE5C365", // 注释的颜色
    // "keywords": "#D15FEE",   // 关键字的颜色
    // "variables": "#eff2f3",  // 变量名的颜色
    // "strings": "#e8c112",    // 字符串的颜色
    // "functions": "#4A4AFF",  // 函数的颜色
    // "types": "#FF0000",      // 类型名的颜色
    // "numbers": "#FF0000",    // 数字的颜色
    // "textMateRules": [],     // 其它具体的颜色设置
  },
  "update.mode": "manual",
  "workbench.iconTheme": "material-icon-theme"
}
```

```json
{
  "editor.quickSuggestionsDelay": 0,
  "editor.rename.enablePreview": false,
  "editor.unicodeHighlight.nonBasicASCII": false,
  "remote.SSH.showLoginTerminal": true,
  "remote.downloadExtensionsLocally": true,
  "remote.SSH.defaultForwardedPorts": [],
  "remote.SSH.remotePlatform": {
    "ip": "linux"
  },
  "remote.SSH.configFile": "C:\\Users\\admin\\.ssh\\config",
  "workbench.colorCustomizations": {
    //覆盖当前所选颜色主题的颜色
    "editor.background": "#241d1d", //编辑器背景色
    "editor.selectionBackground": "#6b5244", //用户选中代码段的颜色
    "editor.findMatchBackground": "#ff0000", //当前搜索匹配的颜色
    "editor.findMatchHighlightBackground": "#ff00ff", //其他搜索匹配项的颜色
    "editor.findRangeHighlightBackground": "#ff9900", //限制搜索范围的颜色
    "editor.lineHighlightBackground": "#48314e", //光标所在行高亮内容的背景颜色
    "editor.lineHighlightBorder": "#704b36" //光标所在行四周边框的背景颜色
  },
  "workbench.preferredDarkColorTheme": "Monokai",
  "workbench.preferredHighContrastColorTheme": "Abyss",
  "workbench.preferredLightColorTheme": "Monokai",
  "workbench.colorTheme": "Monokai",
  "workbench.editor.enablePreviewFromQuickOpen": true,
  "workbench.editor.enablePreview": false,
  "workbench.startupEditor": "none",
  "window.zoomLevel": 1,
  "window.openFoldersInNewWindow": "on",
  "window.newWindowDimensions": "offset",
  "security.workspace.trust.untrustedFiles": "open",
  "[jsonc]": {
    "editor.quickSuggestions": {
      "strings": true
    },
    "editor.suggest.insertMode": "replace"
  },
  "json.schemaDownload.enable": false,
  "files.associations": {
    "*.json": "jsonc"
  },
  "files.autoSave": "onFocusChange",
  "explorer.confirmDelete": false,
  "update.mode": "manual",
  "update.enableWindowsBackgroundUpdates": false,
  "search.followSymlinks": false,
  "go.useLanguageServer": false,
  "[go]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[go.mod]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "python.languageServer": "Default",
  "python.linting.enabled": true,
  "python.formatting.provider": "yapf",
  "[python]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.defaultFormatter": "ms-python.python"
  },
  "python.formatting.yapfArgs": [
    "--style={column_limit=128}",
    "max-line-length:120" // 解决python代码自动换行问题
  ],
  "python.linting.pylintEnabled": false,
  "git.openRepositoryInParentFolders": "always",
  "editor.unicodeHighlight.allowedLocales": {
    "zh-hant": true
  },
  "extensions.ignoreRecommendations": true,
  "files.exclude": {
    "**/.*": false
  },
  "C_Cpp.intelliSenseEngineFallback": "disabled",
  "C_Cpp.intelliSenseEngine": "Tag Parser"
}
```

```json
{
  "editor.fontSize": 16,
  "editor.formatOnPaste": true,
  "editor.fontLigatures": false,
  "editor.formatOnSave": true,
  "editor.formatOnType": true,
  "editor.links": false,
  "editor.fontWeight": "normal",
  "editor.codeActionsOnSave": {
    "source.organizeImports": true
    // "source.fixAll": true  // 注释掉不然保存很卡
  },
  "editor.quickSuggestionsDelay": 0,
  "editor.rename.enablePreview": false,
  "editor.unicodeHighlight.nonBasicASCII": false,
  "editor.tokenColorCustomizations": {
    // "comments": "#FFE5C365", // 注释的颜色
    // "keywords": "#D15FEE",   // 关键字的颜色
    // "variables": "#eff2f3",  // 变量名的颜色
    // "strings": "#e8c112",    // 字符串的颜色
    // "functions": "#4A4AFF",  // 函数的颜色
    // "types": "#FF0000",      // 类型名的颜色
    // "numbers": "#FF0000",    // 数字的颜色
    // "textMateRules": [],     // 其它具体的颜色设置
  },
  "editor.unicodeHighlight.allowedLocales": {
    "zh-hant": true
  },
  "remote.SSH.showLoginTerminal": true,
  "remote.downloadExtensionsLocally": true,
  "remote.SSH.defaultForwardedPorts": [],
  "remote.SSH.remotePlatform": {
    "ip": "linux"
  },
  "remote.SSH.configFile": "C:\\Users\\admin\\.ssh\\config",
  "workbench.colorCustomizations": {
    //覆盖当前所选颜色主题的颜色
    "editor.background": "#241d1d", //编辑器背景色
    "editor.selectionBackground": "#6b5244", //用户选中代码段的颜色
    "editor.findMatchBackground": "#ff0000", //当前搜索匹配的颜色
    "editor.findMatchHighlightBackground": "#ff00ff", //其他搜索匹配项的颜色
    "editor.findRangeHighlightBackground": "#ff9900", //限制搜索范围的颜色
    "editor.lineHighlightBackground": "#48314e", //光标所在行高亮内容的背景颜色
    "editor.lineHighlightBorder": "#704b36" //光标所在行四周边框的背景颜色
  },
  "workbench.preferredDarkColorTheme": "Monokai",
  "workbench.preferredHighContrastColorTheme": "Abyss",
  "workbench.preferredLightColorTheme": "Monokai",
  "workbench.editor.enablePreviewFromQuickOpen": true,
  "workbench.startupEditor": "none",
  "workbench.editor.enablePreview": false,
  "workbench.iconTheme": "material-icon-theme",
  "window.openFoldersInNewWindow": "on",
  "window.newWindowDimensions": "offset",
  "window.zoomLevel": 1,
  "security.workspace.trust.untrustedFiles": "open",
  "[jsonc]": {
    "editor.quickSuggestions": {
      "strings": true
    },
    "editor.suggest.insertMode": "replace"
  },
  "json.schemaDownload.enable": false,
  "files.associations": {
    "*.json": "jsonc"
  },
  "files.autoSave": "afterDelay",
  "explorer.confirmDelete": false,
  "explorer.openEditors.minVisible": 1,
  "search.followSymlinks": false,
  "[go]": {
    "editor.snippetSuggestions": "none",
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "gopls": {
    "usePlaceholders": true, // add parameter placeholders when completing a function
    // Experimental settings
    "completeUnimported": true, // autocomplete unimported packages
    "watchFileChanges": true, // watch file changes outside of the editor
    "deepCompletion": true // enable deep completion
  },
  "[go.mod]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "go.useLanguageServer": true,
  "go.inferGopath": false,
  "files.eol": "\n",
  "python.languageServer": "Default",
  "[python]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.defaultFormatter": "ms-python.black-formatter"
  },
  "git.openRepositoryInParentFolders": "always",
  "extensions.ignoreRecommendations": true,
  "C_Cpp.intelliSenseEngineFallback": "disabled",
  "C_Cpp.intelliSenseEngine": "Tag Parser",
  "update.mode": "manual",
  "Codegeex.Privacy": true,
  "cmake.configureOnOpen": false
}
```

```json
{
  "editor.fontSize": 16,
  "editor.formatOnPaste": true,
  "editor.fontLigatures": false,
  "editor.formatOnSave": true,
  "editor.formatOnType": true,
  "editor.links": false,
  "editor.fontWeight": "normal",
  "editor.codeActionsOnSave": {
    // "source.organizeImports": true,
    // "source.fixAll": true
  },
  "editor.quickSuggestionsDelay": 0,
  "editor.rename.enablePreview": false,
  "editor.unicodeHighlight.nonBasicASCII": false,
  "editor.tokenColorCustomizations": {
    // "comments": "#FFE5C365", // 注释的颜色
    // "keywords": "#D15FEE",   // 关键字的颜色
    // "variables": "#eff2f3",  // 变量名的颜色
    // "strings": "#e8c112",    // 字符串的颜色
    // "functions": "#4A4AFF",  // 函数的颜色
    // "types": "#FF0000",      // 类型名的颜色
    // "numbers": "#FF0000",    // 数字的颜色
    // "textMateRules": [],     // 其它具体的颜色设置
  },
  "editor.unicodeHighlight.allowedLocales": {
    "zh-hant": true
  },
  "remote.SSH.showLoginTerminal": true,
  "remote.downloadExtensionsLocally": true,
  "remote.SSH.defaultForwardedPorts": [],
  "remote.SSH.remotePlatform": {
    "ip": "linux"
  },
  "remote.SSH.configFile": "C:\\Users\\admin\\.ssh\\config",
  "workbench.colorCustomizations": {
    //覆盖当前所选颜色主题的颜色
    "editor.background": "#241d1d", //编辑器背景色
    "editor.selectionBackground": "#6b5244", //用户选中代码段的颜色
    "editor.findMatchBackground": "#ff0000", //当前搜索匹配的颜色
    "editor.findMatchHighlightBackground": "#ff00ff", //其他搜索匹配项的颜色
    "editor.findRangeHighlightBackground": "#ff9900", //限制搜索范围的颜色
    "editor.lineHighlightBackground": "#48314e", //光标所在行高亮内容的背景颜色
    "editor.lineHighlightBorder": "#704b36" //光标所在行四周边框的背景颜色
  },
  "workbench.preferredDarkColorTheme": "Monokai",
  "workbench.preferredHighContrastColorTheme": "Abyss",
  "workbench.preferredLightColorTheme": "Monokai",
  "workbench.editor.enablePreviewFromQuickOpen": true,
  "workbench.startupEditor": "none",
  "workbench.editor.enablePreview": false,
  "workbench.iconTheme": "material-icon-theme",
  "window.openFoldersInNewWindow": "on",
  "window.newWindowDimensions": "offset",
  "window.zoomLevel": 1,
  "security.workspace.trust.untrustedFiles": "open",
  "[jsonc]": {
    "editor.quickSuggestions": {
      "strings": true
    },
    "editor.suggest.insertMode": "replace",
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "json.schemaDownload.enable": false,
  "files.associations": {
    "*.json": "jsonc"
  },
  "files.autoSave": "afterDelay",
  "explorer.confirmDelete": false,
  "explorer.openEditors.minVisible": 1,
  "search.followSymlinks": false,
  "[go]": {
    // "editor.snippetSuggestions": "none",
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.defaultFormatter": "golang.go"
  },
  "gopls": {
    "usePlaceholders": true, // add parameter placeholders when completing a function
    // Experimental settings
    "completeUnimported": true, // autocomplete unimported packages
    // "watchFileChanges": true, // watch file changes outside of the editor
    "deepCompletion": true // enable deep completion
  },
  "[go.mod]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "go.useLanguageServer": true,
  "go.inferGopath": false,
  "files.eol": "\n",
  "python.languageServer": "Default",
  "[python]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.defaultFormatter": "ms-python.black-formatter" // 格式化"[该语言代码]"
  },
  "git.openRepositoryInParentFolders": "always",
  "extensions.ignoreRecommendations": true,
  "C_Cpp.intelliSenseEngineFallback": "disabled",
  "C_Cpp.intelliSenseEngine": "Tag Parser",
  "update.mode": "manual",
  "Codegeex.Privacy": true,
  "cmake.configureOnOpen": false,
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[html]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[css]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[less]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[vue]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "prettier.printWidth": 160
}
```

```json
{
  "editor.fontSize": 16,
  "editor.formatOnPaste": true,
  "editor.fontLigatures": false,
  "editor.formatOnType": true,
  "editor.links": false,
  "editor.fontWeight": "normal",
  "editor.quickSuggestionsDelay": 0,
  "editor.rename.enablePreview": false,
  "editor.unicodeHighlight.nonBasicASCII": false,
  "editor.codeActionsOnSave": {
    // "source.organizeImports": true,
    // "source.fixAll": true
  },
  "editor.tokenColorCustomizations": {
    // "comments": "#FFE5C365", // 注释的颜色
    // "keywords": "#D15FEE",   // 关键字的颜色
    // "variables": "#eff2f3",  // 变量名的颜色
    // "strings": "#e8c112",    // 字符串的颜色
    // "functions": "#4A4AFF",  // 函数的颜色
    // "types": "#FF0000",      // 类型名的颜色
    // "numbers": "#FF0000",    // 数字的颜色
    // "textMateRules": [],     // 其它具体的颜色设置
  },
  "editor.unicodeHighlight.allowedLocales": {
    "zh-hant": true
  },
  "editor.wordWrap": "on",
  "remote.SSH.showLoginTerminal": true,
  "remote.downloadExtensionsLocally": true,
  "remote.SSH.defaultForwardedPorts": [],
  "remote.SSH.remotePlatform": {
    "ip": "linux",
    "ip2": "linux",
  },
  "remote.SSH.configFile": "C:\\Users\\admin\\.ssh\\config",
  "workbench.colorCustomizations": {
    //覆盖当前所选颜色主题的颜色
    "editor.background": "#241d1d", //编辑器背景色
    "editor.selectionBackground": "#6b5244", //用户选中代码段的颜色
    "editor.findMatchBackground": "#ff0000", //当前搜索匹配的颜色
    "editor.findMatchHighlightBackground": "#ff00ff", //其他搜索匹配项的颜色
    "editor.findRangeHighlightBackground": "#ff9900", //限制搜索范围的颜色
    "editor.lineHighlightBackground": "#48314e", //光标所在行高亮内容的背景颜色
    "editor.lineHighlightBorder": "#704b36" //光标所在行四周边框的背景颜色
  },
  "workbench.preferredDarkColorTheme": "Monokai",
  "workbench.preferredHighContrastColorTheme": "Abyss",
  "workbench.preferredLightColorTheme": "Monokai",
  "workbench.editor.enablePreviewFromQuickOpen": true,
  "workbench.startupEditor": "none",
  "workbench.iconTheme": "material-icon-theme",
  "workbench.editor.enablePreview": false,
  "window.openFoldersInNewWindow": "on",
  "window.newWindowDimensions": "offset",
  "window.zoomLevel": 1,
  "security.workspace.trust.untrustedFiles": "open",
  "breadcrumbs.enabled": true, // 开启 vscode 文件路径导航
  "explorer.confirmDelete": false,
  "explorer.openEditors.minVisible": 1,
  "search.followSymlinks": false,
  "files.eol": "\n",
  "black-formatter.args": ["--line-length", "1000"],
  "git.openRepositoryInParentFolders": "always",
  "extensions.ignoreRecommendations": true,
  "update.mode": "manual",
  "[go]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.defaultFormatter": "golang.go"
    // "editor.snippetSuggestions": "none",
  },
  "gopls": {
    "usePlaceholders": true, // add parameter placeholders when completing a function Experimental settings
    "completeUnimported": true, // autocomplete unimported packages
    "deepCompletion": true // enable deep completion
    // "watchFileChanges": true, // watch file changes outside of the editor
  },
  "[go.mod]": {
    "editor.formatOnSave": true,
    "editor.formatOnType": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "go.vetOnSave": "off",
  "go.useLanguageServer": true,
  "go.inferGopath": false,
  "go.formatTool": "goimports",
  "go.languageServerFlags": ["serve", "-rpc.trace", "--debug=localhost:6060"],
  "[go][go.mod][python]": {
    "editor.codeActionsOnSave": {
      "source.organizeImports": "explicit"
    }
  },
  "go.toolsManagement.autoUpdate": true,
  "python.languageServer": "Default",
  "[python]": {
    "editor.formatOnType": true,
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "ms-python.black-formatter", // 格式化"[该语言代码]"
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
    // "editor.wordWrap": "wordWrapColumn",
    // "editor.wordWrapColumn": 1000,
  },
  "[javascript]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[html]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[css]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[less]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[vue]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[jsonc]": {
    "editor.quickSuggestions": {
      "strings": true
    },
    "editor.wordWrap": "on",
    "editor.wordWrapColumn": 300,
    "editor.formatOnSave": true,
    "editor.suggest.insertMode": "replace",
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "files.associations": {
    "*.json": "jsonc"
  },
  "json.schemaDownload.enable": false,
  "C_Cpp.intelliSenseEngineFallback": "disabled",
  "C_Cpp.intelliSenseEngine": "Tag Parser",
  "cmake.configureOnOpen": false,
  "cmake.options.advanced": {
    "build": {
      "statusBarVisibility": "inherit",
      "inheritDefault": "visible"
    },
    "launch": {
      "statusBarVisibility": "inherit",
      "inheritDefault": "visible"
    },
    "debug": {
      "statusBarVisibility": "inherit",
      "inheritDefault": "visible"
    }
  },
  "terminal.integrated.enableMultiLinePasteWarning": false,
  "prettier.printWidth": 160,
  "liveServer.settings.donotShowInfoMsg": true,
  "settingsSync.ignoredExtensions": ["cweijan.vscode-office"],
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "workbench.editorAssociations": {
    "{git,gitlens,git-graph}:/**/*.{md,csv,svg}": "default"
  }
}

```


### Go 的代码片段

```json
{
	// Place your snippets for go here. Each snippet is defined under a snippet name and has a prefix, body and
	// description. The prefix is what is used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
	// $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders. Placeholders with the
	// same ids are connected.
	// Example:
	// "Print to console": {
	// 	"prefix": "log",
	// 	"body": [
	// 		"console.log('$1');",
	// 		"$2"
	// 	],
	// 	"description": "Log output to console"
	// }
		"HEADER": {
			"prefix": "header",
			"body": [
				"/*",
				"@File   : $TM_FILENAME",
				"@Author : pan",
				"@Time   : $CURRENT_YEAR-$CURRENT_MONTH-$CURRENT_DATE $CURRENT_HOUR:$CURRENT_MINUTE:$CURRENT_SECOND",
				"*/",
				"",
			]
		}
}
```

### Python 的代码片段

```json
{
	"HEADER":{
		"prefix": "header",
		"body": [
			"# -*- encoding: utf-8 -*-",
			"'''",
			"@File   : $TM_FILENAME",
			"@Time   : $CURRENT_YEAR-$CURRENT_MONTH-$CURRENT_DATE $CURRENT_HOUR:$CURRENT_MINUTE:$CURRENT_SECOND",
			"@Author : pan",
			"'''",
			"",
		]
	}
}
```

### C 的代码片段

```json
{
	// Place your snippets for c here. Each snippet is defined under a snippet name and has a prefix, body and
	// description. The prefix is what is used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
	// $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders. Placeholders with the
	// same ids are connected.
	// Example:
	// "Print to console": {
	// 	"prefix": "log",
	// 	"body": [
	// 		"console.log('$1');",
	// 		"$2"
	// 	],
	// 	"description": "Log output to console"
	// }
	"HEADER": {
		"prefix": "header",
		"body": [
			"/*",
			"@File   : $TM_FILENAME",
			"@Author : pan",
			"@Time   : $CURRENT_YEAR-$CURRENT_MONTH-$CURRENT_DATE $CURRENT_HOUR:$CURRENT_MINUTE:$CURRENT_SECOND",
			"*/",
			"",
		]
	}
}
```
