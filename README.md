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