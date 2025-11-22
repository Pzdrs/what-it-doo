# What It Do

*a WhatsApp alternative brought to you by the hood*

**What It Do** is a real-time distributed messaging application ticking all the boxes of a modern chat app.

## Demo

Access the demo at [wid.pycrs.cz](https://wid.pycrs.cz).

## Features

- [x] Typing indicators
- [x] Cloud-native
- [x] Sending and sent indicators
- [ ] Delivered and read receipts
- [x] Group chats
- [ ] Configurable 3rd party IdP
- [x] Gravatar support
- [ ] Reply to message

## Architecture

### Infrastructure

The application is designed to run on Kubernetes, leveraging its scalability and resilience features. It uses PostgreSQL as the primary database for storing user data and messages and Redis for inter-service communication.

Both workers and the API servers can be independently scaled up and down and all communication is handled via the Redis-based bus.

[![](https://mermaid.ink/img/pako:eNqVk21rgzAQx79KuFct1FKNdhrGYA9vBh1sXWGw2hepuVnZNCXGPZV-90XTbkq7wXwT7_6_-3t3mA0kUiAw4CpZZRoTXSl0lqh5XBDzpEpWa8LXWS95kZXoz89vr0mJ6hVVuWgjb1I9m9wee7ChQSxUl2QJEoUiK3uCa77kJfbn0zpedBmxbAG3stSpwvu7yYGXbcPt2bM_v2_OBcmKuuFjrPcPlv7Kdmk7-E8XdvKG3u3kGO_9k6d_8j8VphM2I6eOc0YuWLPstub9odEDres6IY5Dpkwsu4ZH07SdtsJuTeyiFmbf_G4bx9O0ne7aTG2vk_Yce6_fNXqgxQUMIFWZAKZVhQPIUeW8DmFTV8agV5hjDMy8Cq6eY4iLralZ8-JRynxfZu5AugL2xF9KE1Vr8wPjVcZTxfPvrMJCoLqUVaGBudQfNS7ANvAOzPfdYRh5ozCMxmN_7LnBAD6AeZQOo8APoiCITsKRG0TbAXw2Hx4NwxMDmTm0VDf2FjeXefsFXYY-pg?type=png)](https://mermaid.live/edit#pako:eNqVk21rgzAQx79KuFct1FKNdhrGYA9vBh1sXWGw2hepuVnZNCXGPZV-90XTbkq7wXwT7_6_-3t3mA0kUiAw4CpZZRoTXSl0lqh5XBDzpEpWa8LXWS95kZXoz89vr0mJ6hVVuWgjb1I9m9wee7ChQSxUl2QJEoUiK3uCa77kJfbn0zpedBmxbAG3stSpwvu7yYGXbcPt2bM_v2_OBcmKuuFjrPcPlv7Kdmk7-E8XdvKG3u3kGO_9k6d_8j8VphM2I6eOc0YuWLPstub9odEDres6IY5Dpkwsu4ZH07SdtsJuTeyiFmbf_G4bx9O0ne7aTG2vk_Yce6_fNXqgxQUMIFWZAKZVhQPIUeW8DmFTV8agV5hjDMy8Cq6eY4iLralZ8-JRynxfZu5AugL2xF9KE1Vr8wPjVcZTxfPvrMJCoLqUVaGBudQfNS7ANvAOzPfdYRh5ozCMxmN_7LnBAD6AeZQOo8APoiCITsKRG0TbAXw2Hx4NwxMDmTm0VDf2FjeXefsFXYY-pg)

### The components

#### API Server

The API server acts as the RESTful interface for clients to interact with the application. It handles user authentication, message sending and retrieval, and other core functionalities.

It also acts as a WebSocket gateway, allowing real-time communication between clients and the server.

#### Workers

In the name of scalability, most of the processing invoked by an incoming WebSocket message is offloaded to a pool of worker services so the server can just fire-and-forget and be ready to handle the next incoming message.

Pretty much the only task handled by workers in this stage of the application is processing and storing incoming messages.

#### PostgreSQL

The primary data store for the application, PostgreSQL is used to store user information, messages, and other relevant data. Thats it.

#### Redis

I'm not doing any caching whatsoever, so Redis is used solely as a communication layer between the API servers and the workers.

### The communication bus

All communication between the API servers and the workers is done via a Redis-based message bus. This allows for decoupling of services and enables easy scaling of the worker pool.

There are two methods of communication:

- Tasks
  - Workers are always on the receiving end of tasks dispatched by the API servers
  - **Redis streams** are used to implement tasks
- Events
  - Events can be dispatched by both API servers and workers but are always received by API servers
  - We use a single *global channel* subscribed to by all API servers for things like user presence updates and message fan-outs. For events targeting specific users, we use *per-gateway channels* to avoid unnecessary message processing
  - **Redis Pub/Sub** is used to implement events

#### Flow examples

##### Sending a message

[![](https://mermaid.ink/img/pako:eNp1k19r2zAUxb_KRS9rIDPx3yR6KHQthDEGJiUEhl9U-8YRtqVMkpN1Id9913E9F9L6xb7S79xzJFlnlusCGWcWf7eocnySojSiyRTQcxDGyVwehHKwAWFhY9HAN6NP9L5FVh2yxZdnnVdIpXB4Eq9w95B-h2c0RzSTW9G6E62xkPZ2bnttqE31kVvaTabautIgaXtg8_X-fsXB98hQFdCgtaJEOEoxBuvJFZFrDoEHaftSS7sHJ2wFTr8PsyZoyyHsoLruiZ3RzXtmS0zKISJLpw2C2-PgOwLkFHvwJO1BuHwPD48_AI-o3KfEkHwnlG7dmIYWl3hkn6M84tgH7iytF80XC2W_7ZPPNH3HQaYprhk0djLuzYbDvFu2HbOIvFL6VGNRYtOJP4SlynUjVTmo2JSVRhaMO9PilDVoGtGV7NzJM0b-DWaM02chTJWxTF1IQ0f8S-tmkBndlnvGd6K2VLWHggK__ar_R811Bx51qxzjYRzF1y6Mn9kfxv1g6fmxn8RxNPODRRjQ7Cvjy9Dzk2geRuHMD5NkvrxM2d-r78xbJGHgz5PYX0TRLIjmU0aHTkf8s78x14tz-Qfr8wX8?type=png)](https://mermaid.live/edit#pako:eNp1k19r2zAUxb_KRS9rIDPx3yR6KHQthDEGJiUEhl9U-8YRtqVMkpN1Id9913E9F9L6xb7S79xzJFlnlusCGWcWf7eocnySojSiyRTQcxDGyVwehHKwAWFhY9HAN6NP9L5FVh2yxZdnnVdIpXB4Eq9w95B-h2c0RzSTW9G6E62xkPZ2bnttqE31kVvaTabautIgaXtg8_X-fsXB98hQFdCgtaJEOEoxBuvJFZFrDoEHaftSS7sHJ2wFTr8PsyZoyyHsoLruiZ3RzXtmS0zKISJLpw2C2-PgOwLkFHvwJO1BuHwPD48_AI-o3KfEkHwnlG7dmIYWl3hkn6M84tgH7iytF80XC2W_7ZPPNH3HQaYprhk0djLuzYbDvFu2HbOIvFL6VGNRYtOJP4SlynUjVTmo2JSVRhaMO9PilDVoGtGV7NzJM0b-DWaM02chTJWxTF1IQ0f8S-tmkBndlnvGd6K2VLWHggK__ar_R811Bx51qxzjYRzF1y6Mn9kfxv1g6fmxn8RxNPODRRjQ7Cvjy9Dzk2geRuHMD5NkvrxM2d-r78xbJGHgz5PYX0TRLIjmU0aHTkf8s78x14tz-Qfr8wX8)