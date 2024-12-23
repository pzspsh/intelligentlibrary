# -*- encoding: utf-8 -*-
'''
@File   : main.py
@Time   : 2023-11-08 10:02:53
@Author : pan
'''
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import padding
import base64


def aes_decrypt(key, iv, ciphertext):
    # 解码密钥和初始向量
    key = base64.b64decode(key)
    iv = base64.b64decode(iv)

    # 创建解密器
    backend = default_backend()
    cipher = Cipher(algorithms.AES(key), modes.CBC(iv), backend=backend)
    decryptor = cipher.decryptor()

    # 解密密文
    padded_ciphertext = base64.b64decode(ciphertext)
    plaintext = decryptor.update(padded_ciphertext) + decryptor.finalize()

    # 去除填充
    unpadder = padding.PKCS7(128).unpadder()
    plaintext = unpadder.update(plaintext) + unpadder.finalize()

    return plaintext.decode("utf-8")


# 示例参数
key = "L3Vzci9sb2NhbC9iYXNlNjQva2V5LWJpbmFyeQ=="
iv = "czZCaGRSa3F0MzpnWDFmQmF0M2JW"
ciphertext = "CJ0tBdP4zUl0r7zU0MSnIw=="

# 解密密文
plaintext = aes_decrypt(key, iv, ciphertext)
print("明文：", plaintext)