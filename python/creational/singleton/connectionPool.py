class ConnectionPool(object):
    _instance = None 

    def __new__(cls):
        if cls._instance is None:
            print("Creating the object")
            cls._instance = super(ConnectionPool, cls).__new__(cls)
        return cls._instance

pool1 = ConnectionPool()
print(pool1)

pool2 = ConnectionPool()
print(pool2)
print("Are they the same object?", pool1 is pool2)