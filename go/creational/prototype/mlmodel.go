package prototype

type ObjectClonable interface {
	Clone() *MLModel
}

type MLModel struct {
	modelType       ModelType
	description     string
	trainingSplit   float64
	validationSplit float64
	alpha           float64
	beta            float64
}

func (m *MLModel) Clone() *MLModel {
	return &MLModel{
		modelType:       m.modelType,
		description:     m.description,
		trainingSplit:   m.trainingSplit,
		validationSplit: m.validationSplit,
		alpha:           m.alpha,
		beta:            m.beta,
	}
}

// Getter and Setter methods
func (m *MLModel) GetModelType() ModelType {
	return m.modelType
}

func (m *MLModel) SetModelType(modelType ModelType) {
	m.modelType = modelType
}

func (m *MLModel) GetDescription() string {
	return m.description
}

func (m *MLModel) SetDescription(description string) {
	m.description = description
}

func (m *MLModel) GetTrainingSplit() float64 {
	return m.trainingSplit
}

func (m *MLModel) SetTrainingSplit(trainingSplit float64) {
	m.trainingSplit = trainingSplit
}

func (m *MLModel) GetValidationSplit() float64 {
	return m.validationSplit
}

func (m *MLModel) SetValidationSplit(validationSplit float64) {
	m.validationSplit = validationSplit
}

func (m *MLModel) GetAlpha() float64 {
	return m.alpha
}

func (m *MLModel) SetAlpha(alpha float64) {
	m.alpha = alpha
}

func (m *MLModel) GetBeta() float64 {
	return m.beta
}

func (m *MLModel) SetBeta(beta float64) {
	m.beta = beta
}

type ModelType string

const (
	LR  ModelType = "LR"
	SVM ModelType = "SVM"
	DT  ModelType = "DT"
)

type ModelRegistry struct {
	models map[ModelType]*MLModel
}

func NewModelRegistry() *ModelRegistry {
	return &ModelRegistry{
		models: make(map[ModelType]*MLModel),
	}
}

func (r *ModelRegistry) RegisterModel(model *MLModel) {
	r.models[model.GetModelType()] = model
}

func (r *ModelRegistry) GetModel(modelType ModelType) *MLModel {
	if model, exists := r.models[modelType]; exists {
		return model.Clone()
	}
	return nil
}
