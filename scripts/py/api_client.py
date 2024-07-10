import requests
import hmac
import hashlib
import time
import uuid

class Client:
    address = ""
    def __init__(self):
        self.address = "http://localhost:8080"

    def generate_signature(self, timestamp, nonce):
        h = hmac.new(self.seckey.encode(), digestmod=hashlib.sha256)
        h.update(self.appid.encode())
        h.update(self.apikey.encode())
        h.update(timestamp.encode())
        h.update(nonce.encode())
        return h.hexdigest()

    def request(self, method, urlPath, body=None, headers=None):
        headers = headers or {}
        timestamp = str(int(time.time()))
        nonce = str(uuid.uuid1())
        headers["X-Timestamp"] = timestamp
        headers["X-Nonce"] = nonce
        headers["Content-Type"] = "application/json"
        print("Requesting: " + self.address + urlPath + " with method: " + method)
        print("Body: " + str(body))
        print("Headers: " + str(headers))
        url = self.address + urlPath
        if method == "GET":
            return requests.get(url, headers=headers)
        elif method == "POST":
            return requests.post(url, data=body, headers=headers)
        
    def post(self, urlPath, body):
        return self.request("POST", urlPath, body)
    
    def get(self, urlPath):
        return self.request("GET", urlPath)