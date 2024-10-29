# -*- encoding: utf-8 -*-
"""
@File   : update.py
@Time   : 2023-06-02 13:32:49
@Author : pan
"""
from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

# 定义模型
Base = declarative_base()


class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True)
    name = Column(String)
    age = Column(Integer)


def update():
    # 连接数据库
    username = ""
    password = ""
    host = "127.0.0.1"
    port = 5432
    db = "数据库名"
    engine = create_engine(f"postgresql://{username}:{password}@{host}:{port}/{db}")
    Base.metadata.create_all(engine)  # 创建表users
    Session = sessionmaker(bind=engine)
    session = Session()

    # 更新操作：将名字为'John Doe'的用户的年龄更新为30
    session.query(User).filter(User.name == "John Doe").update({User.age: 30})
    session.commit()


def update2():
    username = ""
    password = ""
    host = "127.0.0.1"
    port = 5432
    db = "数据库名"
    engine = create_engine(f"postgresql://{username}:{password}@{host}:{port}/{db}")
    Base.metadata.create_all(engine)
    Session = sessionmaker(bind=engine)
    session = Session()

    # 假设我们要修改name为'Alice'的用户的age为30
    user_to_update = session.query(User).filter_by(name="Alice").first()
    if user_to_update:
        user_to_update.age = 30
        session.commit()

    # 查看修改后的结果
    updated_users = session.query(User).filter_by(name="Alice").all()
    for user in updated_users:
        print(user.age)  # 应该打印出30


if __name__ == "__main__":
    update()
    pass
