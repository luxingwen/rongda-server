import api_client
import json

class UserReq:
    def __init__(self):
        self.client = api_client.Client()

    def get_user(self, user_id):
        user = {
            "uuid": user_id
        }
        r = self.client.post(f'/api/v1/user/info', json.dumps(user))
        print(r.text)
        return r

    def list_users(self):
        user = {
            "pageSize": 10,
            "current":0
        }
        r = self.client.post('/api/v1/user/list', json.dumps(user))
        print(r.text)
        return r

    def create_user(self, user = None):
        if user is None:
            user = {
                "age": 0,
                "avatar": "",
                "email": "lxw@qq.com",
                "id": 0,
                "nickname": "user1",
                "password": "123",
                "phone": "12345657",
                "sex": "1",
                "signed": "1234",
                "status": 0,
                "username": "lxwtest",
            }
        r = self.client.post('/api/v1/user/create', json.dumps(user))   
        print(r.text)
        return r

    def update_user(self, user_id, user):
        return self.client.post(f'/users/{user_id}', user)

    def delete_user(self, user_id):
        return self.client.delete(f'/users/{user_id}')
    
uq = UserReq()
uq.create_user()
uq.get_user("99e6d6e8-a25d-4dff-a17a-936a3aea3d5b")
uq.list_users()