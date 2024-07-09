from abc import ABC, abstractmethod

class Button(ABC):
    def __init__(self, border: float):
        self._border = border

    @property
    def border(self) -> float:
        return self._border

    @border.setter
    def border(self, value: float):
        self._border = value

    @abstractmethod
    def render(self):
        pass

    @abstractmethod
    def onClick(self):
        pass

class ButtonFactory(ABC):
    @abstractmethod
    def createButton(self, border: float, radius: float, length: float) -> Button:
        pass

class RoundButton(Button):
    def __init__(self, border: float, radius: float):
        super().__init__(border)
        self._radius = radius

    @property
    def radius(self) -> float:
        return self._radius

    def onClick(self):
        print("Round Button was clicked!")

    def render(self):
        print("Rendered!")

class RoundButtonFactory(ButtonFactory):
    def createButton(self, border: float, radius: float, length: float) -> Button:
        return RoundButton(border, radius)
    
class SquareButton(Button):
    def __init__(self, border: float, length: float):
        super().__init__(border)
        self._length = length

    @property
    def length(self) -> float:
        return self._length

    def onClick(self):
        print("Square Button was clicked!")

    def render(self):
        print("Rendered!")

class SquareButtonFactory(ButtonFactory):
    def createButton(self, border: float, radius: float, length: float) -> Button:
        return SquareButton(border, length)
    
def client_code(factory: ButtonFactory):
    button = factory.createButton(border=1.0, radius=5.0, length=10.0)

    button.render()
    button.onClick()

def main():
    round_button_factory = RoundButtonFactory()
    print("Using RoundButtonFactory:")
    client_code(round_button_factory)

    square_button_factory = SquareButtonFactory()
    print("\nUsing SquareButtonFactory:")
    client_code(square_button_factory)

if __name__ == "__main__":
    main()