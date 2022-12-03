import threading
import requests
import time

api1 = "https://api.publicapis.org/entries"
api2 = "https://www.boredapi.com/api/activity"

result = {}

def get_api_response(api: str) -> None:

    response = requests.get(api)
    result[api] = response.json()

start = time.time()
get_api_response(api1)
get_api_response(api2)
end = time.time()

print("Time execution sequentialy: ", end - start)

result = {}
start = time.time()
thread_a = threading.Thread(target=get_api_response, args=[api1])
thread_b = threading.Thread(target=get_api_response, args=[api2])
thread_a.start()
thread_b.start()
end = time.time()

thread_a.join()
thread_b.join()
print("Time execution sequentialy: ", end - start)

print(result[api2])


