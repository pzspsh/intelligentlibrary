# -*- encoding: utf-8 -*-
'''
@File   : yamlconfig.py
@Time   : 2023-06-09 17:01:01
@Author : pan
'''
import yaml

def readyaml(filepath):
    try:
        stream = open(filepath, mode="r", encoding="utf-8")
        mode = yaml.load(stream=stream, Loader=yaml.SafeLoader)
        return mode
    except Exception as err:
        return err

def WriteYaml(filepath, data):
    with open(file=filepath, mode="w", encoding="utf-8") as f:
        yaml.dump(data, f)

if __name__ == "__main__":
    filepath = "./pocs/example/test_yaml.yaml"
    result = readyaml(filepath)
    print(result)