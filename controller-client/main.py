from dataclasses import dataclass

import grpc

import cmds_pb2
import cmds_pb2_grpc
import errors


@dataclass
class Agent:
    id: int
    _stub: cmds_pb2_grpc.AgentManagerStub

    def echo(self, text: str) -> str:
        req = cmds_pb2.EchoCommandRequest(agent_id=self.id, text=text)
        resp = self._stub.RunEchoCommand(req)
        return resp.text

    def shell(self, cmd: str, decode: bool = True) -> bytes:
        req = cmds_pb2.ShellCommandRequest(agent_id=self.id, cmd=cmd)
        resp = self._stub.RunShellCommand(req)
        if resp.err:
            raise errors.ShellCommandError(cmd, resp.err)
        return resp.data.decode() if decode else resp.data

    def __repr__(self) -> str:
        return f'Agent(id={self.id})'


@dataclass
class ControllerClient:
    addr: str
    _stub: cmds_pb2_grpc.AgentManagerStub = None

    def connect(self) -> None:
        channel = grpc.insecure_channel(self.addr)
        self._stub = cmds_pb2_grpc.AgentManagerStub(channel)

    def get_agent(self, agent_id: int) -> Agent:
        return Agent(agent_id, self._stub)

    def get_agents(self) -> [cmds_pb2.AgentInfo]:
        return self._stub.GetAgents(cmds_pb2.Empty())


if __name__ == '__main__':
    cnc = ControllerClient('localhost:50052')
    cnc.connect()
