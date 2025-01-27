# Viper

Remote control software using gRPC and Go.

## Components

- Agent - Runs on the endpoint.
- Controller - The server that the _agents_ connect to.
- Agent Manager - The server that runs alongside the _controller_ to manage the _agents_.
- Controller Client - Talks to the _agent manager_.

## Usage

1. Update the `config.json` according to your needs.
2. Run the controller servers: `go run controller/main.go`
3. Run the agents: `go run agent/main.go`
4. Control the agents: `ipython -i controller-client/main.py`

```py
cnc.get_agents()
> [
    id: 0
    alive: true
    connect_time: "2022-06-19 15:22:13.17828 +0300 IDT m=+61.594837928"
]

agent = cnc.get_agent(agent_id=0)
agent.shell('whoami')
> 'root'
```
