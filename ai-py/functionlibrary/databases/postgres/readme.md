# postgres

```python
from loguru import logger
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from configs.configs import Config as cf


def PGConn():
    try:
        databaseUrl = f"postgresql://{cf.pguser}:{cf.pgpass}@{cf.pghost}:{cf.pgport}/{cf.pgdb}"
        engine = create_engine(url=databaseUrl)
        Session = sessionmaker(bind=engine)
        logger.info("Successfully connected to postgres database!")
        return engine, Session()
    except Exception as err:
        logger.error(f"postgress conn error: {err}")
        return None, err


engine, session = PGConn()


from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

Base = declarative_base()

class MyModel(Base):
    __tablename__ = 'my_table'
    id = Column(Integer, primary_key=True)
    field1 = Column(String)
    field2 = Column(String)
    # ... 其他字段
    field10 = Column(String)

# 假设你已经有一个数据库引擎
engine = create_engine('sqlite:///example.db')  # 使用 SQLite 作为示例数据库
Session = sessionmaker(bind=engine)
session = Session()

# 假设你要更新的记录 ID 和新数据
record_id = 1
new_data = {
    'field1': 'new_value1',
    'field2': None,  # 空值，不更新
    'field3': 'new_value3',
    # ... 其他字段
    'field10': None  # 空值，不更新
}

# 查询出需要更新的记录
instance = session.query(MyModel).filter_by(id=record_id).first()

# 如果记录存在，则更新不为空的字段
if instance:
    for field, value in new_data.items():
        if value is not None:  # 检查字段是否为空
            setattr(instance, field, value)  # 更新字段值

    # 提交事务
    session.commit()
else:
    print(f"No record found with id {record_id}")



from sqlalchemy.orm import Session
from your_model_file import YourModel

session = Session(your_engine)
try:
    # 查询出要更新的记录，并更新特定字段
    session.query(YourModel).filter(YourModel.number == 'your_target_number').update(
        {YourModel.field1: 'new_value1', YourModel.field3: 'new_value3'},
        synchronize_session=False  # 可选，用于避免会话同步开销
    )
    session.commit()
except Exception as e:
    session.rollback()
    print(f"An error occurred: {e}")
finally:
    session.close()



from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

# 定义数据库连接（这里以 SQLite 为例）
engine = create_engine('sqlite:///example.db', echo=True)

# 声明基类
Base = declarative_base()

# 定义表模型
class User(Base):
    __tablename__ = 'users'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    age = Column(Integer)

# 创建表
Base.metadata.create_all(engine)

# 创建会话
Session = sessionmaker(bind=engine)
session = Session()

# 要添加或更新的数据
user_data = User(id=1, name='John Doe', age=30)

# 使用 merge 方法
session.merge(user_data)
session.commit()
session.close()


from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

# 定义数据库连接（这里以 SQLite 为例）
engine = create_engine('sqlite:///example.db', echo=True)

# 声明基类
Base = declarative_base()

# 定义表模型
class User(Base):
    __tablename__ = 'users'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    age = Column(Integer)

# 创建表
Base.metadata.create_all(engine)

# 创建会话
Session = sessionmaker(bind=engine)
session = Session()

# 要添加或更新的数据
user_id = 1
user_name = 'John Doe'
user_age = 30

# 查询是否存在
existing_user = session.query(User).filter_by(id=user_id).first()

if existing_user:
    # 更新现有记录
    existing_user.name = user_name
    existing_user.age = user_age
else:
    # 添加新记录
    new_user = User(id=user_id, name=user_name, age=user_age)
    session.add(new_user)

# 提交事务
session.commit()
session.close()






from pony.orm import *
from config.config import Config as cf

db = Database()
db.bind("postgres", user=cf.pguser, password=cf.pgpass, host=cf.pghost, port=cf.pgport, database=cf.pgdb)


class DemoTable(db.Entity):
    _table_ = "Demotable"
    id = PrimaryKey(int, auto=True, column="id")
    number1 = Required(str, max_len=40, column="number1")
    number2 = Required(str, max_len=40, column="number2")

sql_debug(True)
db.generate_mapping(create_tables=True)


@db_session
def insert(number1, number2):
    try:
        DemoTable(number1=number1, number2=number2)
        commit()
    except:
        pass

@db_session
def delete(number1, number2):
    try:
        demo = DemoTable.get(number1=number1, number2=number2)
        demo.delete()
        commit()
    except:
        pass

@db_session
def sesquery(number1, number2):
    try:
        datas = select((p.number, p.number2) for p in DemoTable if p.number1 == number1 and p.number2 == number2)
        if datas:
            return True
        else:
            return False
    except:
        return False

```