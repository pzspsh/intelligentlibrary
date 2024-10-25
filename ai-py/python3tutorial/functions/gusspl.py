# -*- encoding: utf-8 -*-
"""
@File   : guessLanguage.py
@Time   : 2024-10-23 16:58:01
@Author : pan
"""
import re


def GuessPL(code_string):
    """
    Programming Language
    判断代码字符串是那种编程语言
    正则匹配完还有进行对比，除非确定百分比正确
    """
    pypatterns = [  # 待严谨修改 python
        r"\s*def\s+\w+\s*\(.*\)\s*:",
        r"\s*class\s+\w+\s*\(.*\)\s*:",
        r"\s*import\s+(\w+(\s*\.\s*\w+)*)",
        r"\s*from\s+\w+(\s*\.\s*\w+)*\s*import\s+(\w+(\s*,\s*\w+)*)",
        r"\s*return\s+.*",
        r"\s*print\s*\(.*\)",
        r"\s*if\s+.*:",
        r"\s*elif\s+.*:",
        r"\s*else\s*:",
        r"\s*for\s+.*\s*in\s+.*:",
        r"\s*while\s+.*:",
        r"#.*",
    ]
    gopatterns = []  # golang
    cpatterns = []  # c
    cpppatterns = []  # c++
    jspatterns = []  # javascript
    phppatterns = []  # php
    rubypatterns = []  # ruby
    javapatterns = []  # java
    tspatterns = []  # TypeScript
    h5patterns = []  # html
    cssvpatterns = []  # css
    swiffpatterns = []  # swift
    objcpatterns = []  # Objective-C
    sqlpatterns = []  # sql
    rustpatterns = []  # rust
    bashpatterns = []  # bash
    dartpatterns = []  # dart
    rpatterns = []  # R语言
    matlabpatterns = []  # matlab
    luapatterns = []  # lua
    ...
    patterns = [re.compile(pattern, re.MULTILINE) for pattern in pypatterns]
    match_count = sum(len(pattern.findall(code_string)) for pattern in patterns)
    if match_count > 0:
        return "python"
    else:
        return "Unknown Language"


if __name__ == "__main__":
    codestr = """
    """
    GuessPL(codestr)
