from abc import ABC, abstractmethod

class Datasource(ABC):
    @abstractmethod
    def read(self):
        pass

    @abstractmethod
    def write(self, value):
        pass

class BaseDecorator(Datasource):
    def __init__(self, next_layer):
        self._next_layer = next_layer

class FileDatasource(Datasource):
    def read(self):
        return "Base"

    def write(self, value):
        print(value)

class CompressionDecorator(BaseDecorator):
    def __init__(self, datasource):
        super().__init__(datasource)

    def read(self):
        compressed = self._next_layer.read()
        return self._decompress(compressed)

    def _decompress(self, compressed):
        return f"{compressed} - Decompressed"

    def write(self, value):
        compressed = self._compress(value)
        self._next_layer.write(compressed)

    def _compress(self, value):
        return f"{value} - Compressed"

class EncryptionDecorator(BaseDecorator):
    def __init__(self, next_layer):
        super().__init__(next_layer)

    def read(self):
        value = self._next_layer.read()
        return self._decrypt(value)

    def _decrypt(self, value):
        return f"{value} - Decrypted"

    def write(self, value):
        encrypted = self._encrypt(value)
        self._next_layer.write(encrypted)

    def _encrypt(self, value):
        return f"{value} - Encrypted"

if __name__ == "__main__":
    datasource = FileDatasource()
    encrypted_datasource = EncryptionDecorator(datasource)
    compressed_encrypted_datasource = CompressionDecorator(encrypted_datasource)

    compressed_encrypted_datasource.write("Test Data")
    print(compressed_encrypted_datasource.read())
