from __future__ import annotations
from abc import ABC, abstractmethod

class Builder(ABC):
    @abstractmethod
    def with_name(self, name: str) -> Builder:
        pass

    @abstractmethod
    def with_url(self, host: str, port: int) -> Builder:
        pass

    @abstractmethod
    def build(self) -> Database:
        pass

class DatabaseBuilder(Builder):
    def __init__(self) -> None:
        self._database = Database()

    def with_name(self, name: str) -> DatabaseBuilder:
        self._database._name = name
        return self
    
    def with_url(self, host: str, port: int) -> DatabaseBuilder:
        self._database._host = host
        self._database._port = port
        return self
    
    def build(self) -> Database:
        if not self.is_valid():
            raise ValueError("Invalid database configuration")
        return self._database
    
    def is_valid(self) -> bool:
        return self._database._name is not None


class Database:
    def __init__(self, name=None, host=None, port=None):
        self._name = name
        self._host = host
        self._port = port

    def __repr__(self):
        return f"Database(name={self._name}, host={self._host}, port={self._port})"
    
if __name__ == "__main__":
    try:
        # Building a valid database
        db = DatabaseBuilder().with_name("MyDB").with_url("localhost", 3306).build()
        print(db)

        # Building an invalid database (missing name)
        invalid_db = DatabaseBuilder().with_url("localhost", 3306).build()
        print(invalid_db)
    except ValueError as e:
        print(e)