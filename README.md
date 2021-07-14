# raft-sample
A handy raft demo repository of helloworld example from [https://github.com/lni/dragonboat-example](https://github.com/lni/dragonboat-example)

## How it works
Before playing the demo, you must build the binary ```main```,
```
docker compose up build
```

Then to play demo,
```
docker compose up node1 node2 node3
```

Now you'll get
```
node2_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:26.2165754 +0000 UTC m=+7.025916201, count:1
node3_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:26.2165754 +0000 UTC m=+7.025916201, count:1
node1_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:26.2165754 +0000 UTC m=+7.025916201, count:1
node1_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:28.2459868 +0000 UTC m=+9.055463001, count:2
node2_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:28.2459868 +0000 UTC m=+9.055463001, count:2
node3_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:28.2459868 +0000 UTC m=+9.055463001, count:2
node1_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:29.2168722 +0000 UTC m=+10.026242901, count:3
node2_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:29.2168722 +0000 UTC m=+10.026242901, count:3
node3_1  | from ExampleStateMachine.Update(), msg: 2021-07-03 12:59:29.2168722 +0000 UTC m=+10.026242901, count:3
```

After playing demo, you should
```
docker-compose down
```

## Note
In the original example, both of the following are required when adding members!
 1. you have to "add localhost:63100 4" from the node that currently constitutes the cluster
 1. on the new node "./example-helloworld -nodeid 4 -addr localhost:63100 -join" is required on the new node!

That is,
 1.  is not possible to add a node with the same nodeID as a node that has been removed once.
 1. The new nodeID can use the port that has been used once before. Of course, currently used ports cannot be used.
 1. If a process is dropped without "remove", unreachable will be detected on other nodes. When dropping a process, be sure to drop it gracefully.
