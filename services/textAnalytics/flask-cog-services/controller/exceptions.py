

class AnalyticsError(BaseException):

    def __init__(self, message: str, status_code: int) -> None:
        super().__init__(message)
        self.message = message
        self.status_code = status_code

class AzureCognitiveError(AnalyticsError):
    def __init__(self, message: str = "Internal Server Error", status_code: int = 500) -> None:
        super().__init__(message, status_code)