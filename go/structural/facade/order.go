package facade

import "fmt"

// OrderManager interface
type OrderManager interface {
	CreateOrder()
}

// OrderManagerImpl struct
type OrderManagerImpl struct {
	orderProcessor *OrderProcessor
}

func (o *OrderManagerImpl) CreateOrder() {
	o.orderProcessor.Process()
}

// OrderProcessor struct
type OrderProcessor struct {
	recommendationService *RecommendationService
	paymentService        *PaymentService
	warehouseProcessor    *WarehouseProcessor
}

func (o *OrderProcessor) Process() {
	o.warehouseProcessor.Process()
	o.recommendationService.Recommend()
	o.paymentService.Pay()
}

// OrderFlowProcessor struct
type OrderFlowProcessor struct {
	paymentService        *PaymentService
	inventoryService      *InventoryService
	recommendationService *RecommendationService
	analyticsService      *AnalyticsService
}

func NewOrderFlowProcessor() *OrderFlowProcessor {
	return &OrderFlowProcessor{
		paymentService:        &PaymentService{},
		inventoryService:      &InventoryService{},
		recommendationService: &RecommendationService{},
		analyticsService:      &AnalyticsService{},
	}
}

func (o *OrderFlowProcessor) Process() {
	o.paymentService.Pay()
	// update
	o.inventoryService.CheckInventory()
	// analytics
	o.recommendationService.Recommend()
	o.analyticsService.Track()
}

// PaymentService struct
type PaymentService struct{}

func (p *PaymentService) Pay() {
	fmt.Println("Payment done")
}

// RecommendationService struct
type RecommendationService struct{}

func (r *RecommendationService) Recommend() {
	fmt.Println("Recommendation created")
}

// WarehouseProcessor struct
type WarehouseProcessor struct {
	inventoryService *InventoryService
	analyticsService *AnalyticsService
}

func (w *WarehouseProcessor) Process() {
	w.inventoryService.CheckInventory()
	w.analyticsService.Track()
}

// AnalyticsService struct
type AnalyticsService struct{}

func (a *AnalyticsService) Track() {
	fmt.Println("Analytics created")
}

// InventoryService struct
type InventoryService struct{}

func (i *InventoryService) CheckInventory() {
	fmt.Println("Inventory checked")
}
