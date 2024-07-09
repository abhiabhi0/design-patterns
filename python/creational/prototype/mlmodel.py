from abc import ABC, abstractmethod
from copy import deepcopy
from enum import Enum

class ObjectClonable(ABC):
    @abstractmethod
    def clone(self):
        pass

class MLModel(ObjectClonable):
    def __init__(self, model_type, description, training_split, validation_split, alpha, beta):
        self._model_type = model_type
        self._description = description
        self._training_split = training_split
        self._validation_split = validation_split
        self._alpha = alpha
        self._beta = beta

    def clone(self):
        return deepcopy(self)

    # Getter and Setter methods
    @property
    def model_type(self):
        return self._model_type

    @model_type.setter
    def model_type(self, model_type):
        self._model_type = model_type

    @property
    def description(self):
        return self._description

    @description.setter
    def description(self, description):
        self._description = description

    @property
    def training_split(self):
        return self._training_split

    @training_split.setter
    def training_split(self, training_split):
        self._training_split = training_split

    @property
    def validation_split(self):
        return self._validation_split

    @validation_split.setter
    def validation_split(self, validation_split):
        self._validation_split = validation_split

    @property
    def alpha(self):
        return self._alpha

    @alpha.setter
    def alpha(self, alpha):
        self._alpha = alpha

    @property
    def beta(self):
        return self._beta

    @beta.setter
    def beta(self, beta):
        self._beta = beta

    def __str__(self):
        return f'MLModel(type={self._model_type}, description={self._description}, ' \
               f'training_split={self._training_split}, validation_split={self._validation_split}, ' \
               f'alpha={self._alpha}, beta={self._beta})'

class ModelType(Enum):
    LR = "LR"
    SVM = "SVM"
    DT = "DT"

# Step 4: Create and populate registry
class ModelRegistry:
    def __init__(self):
        self._models = {}

    def register_model(self, model):
        self._models[model.model_type] = model

    def get_model(self, model_type):
        model = self._models.get(model_type)
        return model.clone() if model else None

# Example usage
if __name__ == "__main__":
    # Create models
    lr_model = MLModel(ModelType.LR, "Linear Regression Model", 0.7, 0.3, 0.01, 0.1)
    svm_model = MLModel(ModelType.SVM, "Support Vector Machine Model", 0.6, 0.4, 0.02, 0.2)

    # Register models
    registry = ModelRegistry()
    registry.register_model(lr_model)
    registry.register_model(svm_model)

    # Retrieve and clone models
    cloned_lr_model = registry.get_model(ModelType.LR)
    cloned_lr_model.description = "Linear Regression Model Cloned"
    cloned_svm_model = registry.get_model(ModelType.SVM)
    cloned_svm_model.description = "Support Vector Machine Model Cloned"
    print(lr_model)
    print(cloned_lr_model)
    print(svm_model)
    print(cloned_svm_model)
