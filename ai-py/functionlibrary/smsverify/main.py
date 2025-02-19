# -*- encoding: utf-8 -*-
"""
@File   : demo.py
@Time   : 2025-02-18 16:34:22
@Author : pan
"""
from flask import Flask, request, jsonify
from aliyunsdkcore.client import AcsClient
from aliyunsdkcore.request import CommonRequest
import random
import os

app = Flask(__name__)

# 阿里云账户AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
client = AcsClient("<your-access-key-id>", "<your-access-key-secret>", "default")

# 存储验证码的字典（为了简单起见，这里使用内存存储，实际应用中应使用数据库或缓存服务）
verification_codes = {}


@app.route("/get_verification_code", methods=["POST"])
def get_verification_code():
    phone_number = request.json.get("phone_number")
    if not phone_number:
        return jsonify({"error": "Phone number is required"}), 400

    # 生成6位随机验证码
    verification_code = str(random.randint(100000, 999999))

    # 存储验证码（为了简化示例，这里存储在内存中）
    verification_codes[phone_number] = verification_code

    # 设置短信参数
    request = CommonRequest()
    request.set_accept_format("json")
    request.set_domain("dysmsapi.aliyuncs.com")
    request.set_method("POST")
    request.set_protocol_type("https")
    request.set_version("2017-05-25")
    request.set_action_name("SendSms")

    task = {"RegionId": "default", "PhoneNumbers": phone_number, "SignName": "<your-sign-name>", "TemplateCode": "<your-template-code>", "TemplateParam": f'{{"code":{verification_code}}}'}
    request.add_body_params("Task", task)

    # 发送短信
    response = client.do_action_with_exception(request)
    response_data = response.decode("utf-8")

    # 检查发送结果
    if "Code" in response_data and "OK" in response_data:
        return jsonify({"message": "Verification code sent successfully"}), 200
    else:
        return jsonify({"error": "Failed to send verification code", "response": response_data}), 500


if __name__ == "__main__":
    app.run(debug=True)
