# Setting.json

```shell
{
    "editor.fontSize": 16,
    "editor.formatOnPaste": true,
    "editor.fontLigatures": false,
    "editor.links": false,
    "editor.fontWeight": "normal",
    "editor.codeActionsOnSave": {
        "source.organizeImports": true,
        "source.fixAll": true
    },
    "editor.quickSuggestionsDelay": 0,
    "remote.SSH.showLoginTerminal": true,
    "remote.downloadExtensionsLocally": true,
    "remote.SSH.defaultForwardedPorts": [],
    "remote.SSH.remotePlatform": {
        "10.0.35.65": "linux",
        "10.0.35.74": "linux",
        "10.0.35.66": "linux",
        "10.0.35.123": "linux"
    },
    "remote.SSH.configFile": "C:\\Users\\admin\\.ssh\\config",
    "workbench.colorCustomizations": { //覆盖当前所选颜色主题的颜色
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
        "editor.codeActionsOnSave": {
            "source.organizeImports": true
        },
    },
    "[go.mod]": {
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true
        }
    },
    "python.languageServer": "Default",
    "[python]": {
        "editor.formatOnType": true
    },
    "python.linting.flake8Args": [
        "--max--line-length=200"
    ],
    "python.linting.flake8Enabled": true,
    "editor.rename.enablePreview": false,
    "workbench.editor.enablePreviewFromQuickOpen": true,
    "editor.unicodeHighlight.nonBasicASCII": false,
    "workbench.editor.enablePreview": false
}
```