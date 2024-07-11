from abc import ABC, abstractmethod
from enum import Enum

class PaymentStatus(Enum):
    SUCCESS = "SUCCESS"
    FAILURE = "FAILURE"

class PaymentRequest:
    def __init__(self, name, phone, email, amount):
        self._name = name
        self._phone = phone
        self._email = email
        self._amount = amount

    @property
    def name(self):
        return self._name
    
    @property
    def phone(self):
        return self._phone
    
    @property
    def email(self):
        return self._email
    
    @property
    def amount(self):
        return self._amount

class PaymentProviderInterface(ABC):
    @abstractmethod
    def generate_link(self):
        pass
    
    @abstractmethod
    def pay(self, payment_request):
        pass
    
    @abstractmethod
    def check_status(self):
        pass

class CashfreeApi:
    def create_url(self):
        return "Cashfree"
    
    def do_payment(self, amount):
        print(f"Cashfree Payment: {amount}")
    
    def verify_status(self):
        return "OK"

class CashFreePayProvider(PaymentProviderInterface):
    def __init__(self):
        self._cashfree_api = CashfreeApi()
    
    def generate_link(self):
        return self._cashfree_api.create_url()
    
    def pay(self, payment_request):
        self._cashfree_api.do_payment(payment_request.amount)
    
    def check_status(self):
        status = self._cashfree_api.verify_status()
        return self.to_payment_status(status)
    
    def to_payment_status(self, status):
        if status == "OK":
            return PaymentStatus.SUCCESS
        return PaymentStatus.FAILURE

class RazorPayApi:
    def make_link(self):
        return "RazorPay"
    
    def pre_pay(self):
        print("RazorPay PrePayment")
    
    def pay(self, name, amount):
        print(f"RazorPay Payment for {name} of amount {amount}")
    
    def check_status(self):
        # Simulated status check
        return "PASS"

class RazorPayProvider(PaymentProviderInterface):
    def __init__(self):
        self._razorpay_api = RazorPayApi()
    
    def generate_link(self):
        return self._razorpay_api.make_link()
    
    def pay(self, payment_request):
        self._razorpay_api.pre_pay()
        self._razorpay_api.pay(payment_request.name, payment_request.amount)
    
    def check_status(self):
        status = self._razorpay_api.check_status()
        return self.to_payment_status(status)
    
    def to_payment_status(self, status):
        if status == "PASS":
            return PaymentStatus.SUCCESS
        return PaymentStatus.FAILURE

# Client code to execute payment request
def process_payment(payment_provider, payment_request):
    print(payment_provider.generate_link())
    payment_provider.pay(payment_request)
    status = payment_provider.check_status()
    print(f"Payment Status: {status.name}")

# Example usage:
payment_request = PaymentRequest("John Doe", "1234567890", "john@example.com", 100)

# Using CashFreePayProvider
cashfree_provider = CashFreePayProvider()
process_payment(cashfree_provider, payment_request)
