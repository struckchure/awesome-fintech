# awesome-fintech
This project was generated with [Blnk](https://docs.awesome.fintech.org.xyx/white-paper).

# Setup
```sh
$ docker compose up -d
$ go run ./cmd db up
$ go run ./cmd server up
```
If it all checks out, you should have this.
```
[Fx] PROVIDE    *fiber.App <= main.startServer()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             main.ServerRootCmd.func1.1()
[Fx] RUN        provide: go.uber.org/fx.New.func1()
[Fx] RUN        provide: main.startServer()
[Fx] HOOK OnStart               main.startServer.func1() executing (caller: main.startServer)
[Fx] HOOK OnStart               main.startServer.func1() called by main.startServer ran successfully in 4.292µs
[Fx] RUNNING

    _______ __             
   / ____(_) /_  ___  _____
  / /_  / / __ \/ _ \/ ___/
 / __/ / / /_/ /  __/ /    
/_/   /_/_.___/\___/_/          v3.0.0-beta.2
--------------------------------------------------
INFO Server started on:         http://127.0.0.1:3000 (bound on host 0.0.0.0 and port 3000)
INFO Total handlers count:      5
INFO Prefork:                   Disabled
INFO PID:                       51112
INFO Total process count:       1
```
