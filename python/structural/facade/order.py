from abc import ABC, abstractmethod

class OrderManager(ABC):
    @abstractmethod
    def create_order(self) -> None:
        pass

class OrderManagerImpl(OrderManager):
    def __init__(self, order_processor: 'OrderProcessor'):
        self.order_processor = order_processor

    def create_order(self) -> None:
        self.order_processor.process()

class AnalyticsService:
    def track(self) -> None:
        print("Analytics created")

class InventoryService:
    def check_inventory(self) -> None:
        print("Inventory checked")

class OrderProcessor:
    def __init__(self, recommendation_service: 'RecommendationService', 
                 payment_service: 'PaymentService', 
                 warehouse_processor: 'WarehouseProcessor'):
        self.recommendation_service = recommendation_service
        self.payment_service = payment_service
        self.warehouse_processor = warehouse_processor

    def process(self) -> None:
        self.warehouse_processor.process()
        self.recommendation_service.recommend()
        self.payment_service.pay()

class OrderFlowProcessor:
    def __init__(self):
        self.payment_service = PaymentService()
        self.inventory_service = InventoryService()
        self.recommendation_service = RecommendationService()
        self.analytics_service = AnalyticsService()

    def process(self) -> None:
        self.payment_service.pay()
        # update
        self.inventory_service.check_inventory()
        # analytics
        self.recommendation_service.recommend()
        self.analytics_service.track()

class PaymentService:
    def pay(self) -> None:
        print("Payment done")

class RecommendationService:
    def recommend(self) -> None:
        print("Recommendation created")

class WarehouseProcessor:
    def __init__(self, inventory_service: InventoryService, 
                 analytics_service: AnalyticsService):
        self.inventory_service = inventory_service
        self.analytics_service = analytics_service

    def process(self) -> None:
        self.inventory_service.check_inventory()
        self.analytics_service.track()

if __name__ == "__main__":
    order_flow_processor = OrderFlowProcessor()
    order_flow_processor.process()
