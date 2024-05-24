# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2024-05-23 09:56:00
@Author : pan
"""
from fastapi import FastAPI, Request

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}
