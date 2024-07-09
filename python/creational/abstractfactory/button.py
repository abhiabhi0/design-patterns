from abc import ABC, abstractmethod

# Abstract Product: Button
class Button(ABC):
    def __init__(self, border: float):
        self._border = border

    @property
    def border(self) -> float:
        return self._border

    @abstractmethod
    def render(self):
        pass

    @abstractmethod
    def onClick(self):
        pass

# Concrete Product: DarkButton
class DarkButton(Button):
    def __init__(self, border: float, radius: float):
        super().__init__(border)
        self._radius = radius

    @property
    def radius(self) -> float:
        return self._radius

    def onClick(self):
        print("Dark Btn was clicked!")

    def render(self):
        print("Rendered!")

# Abstract Product: Radio
class Radio(ABC):
    @abstractmethod
    def onSelect(self):
        pass

    @abstractmethod
    def render(self):
        pass

# Concrete Product: DarkRadio
class DarkRadio(Radio):
    def onSelect(self):
        print("DarkRadio selected!")

    def render(self):
        print("DarkRadio rendered!")

# Abstract Factory
class ThemeFactory(ABC):
    @abstractmethod
    def createButton(self, border: float, length: float, radius: float) -> Button:
        pass

    @abstractmethod
    def createRadio(self) -> Radio:
        pass

# Concrete Factory: DarkThemeFactory
class DarkThemeFactory(ThemeFactory):
    def createButton(self, border: float, length: float, radius: float) -> Button:
        return DarkButton(border, radius)

    def createRadio(self) -> Radio:
        return DarkRadio()

# Concrete Product: LightButton
class LightButton(Button):
    def __init__(self, border: float, length: float):
        super().__init__(border)
        self._length = length

    @property
    def length(self) -> float:
        return self._length

    def onClick(self):
        print("Light Btn was clicked!")

    def render(self):
        print("Rendered!")

# Concrete Product: LightRadio
class LightRadio(Radio):
    def onSelect(self):
        print("LightRadio selected!")

    def render(self):
        print("LightRadio rendered!")

# Concrete Factory: LightThemeFactory
class LightThemeFactory(ThemeFactory):
    def createButton(self, border: float, length: float, radius: float) -> Button:
        return LightButton(border, length)

    def createRadio(self) -> Radio:
        return LightRadio()

# Client code
def client_code(factory: ThemeFactory):
    button = factory.createButton(border=1.0, length=10.0, radius=5.0)
    button.render()
    button.onClick()

    radio = factory.createRadio()
    radio.render()
    radio.onSelect()

def main():
    print("Using DarkThemeFactory:")
    dark_factory = DarkThemeFactory()
    client_code(dark_factory)

    print("\nUsing LightThemeFactory:")
    light_factory = LightThemeFactory()
    client_code(light_factory)

if __name__ == "__main__":
    main()
