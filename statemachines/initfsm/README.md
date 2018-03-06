## This package defines a rudimentary init FSM

### State Transition Table
| Input\Current State | Idle | Init | Red    | Yellow | Green  |
| ------------------- | ---- | ---- | ------ | ------ | ------ |
| Start               | Init | .... | ...... | ...... | ...... |
| ModuleReady         | .... | Red  | ...... | ...... | ...... |
| CritReady           | .... | .... | Yellow | ...... | ...... |
| CritFail            | .... | .... | Red    | Red    | Red    |
| OptReady            | .... | .... | ...... | Green  | ...... |
| OptFail             | .... | .... | ...... | Yellow | Yellow |

### Digraph
```
digraph fsm {
    ratio = fill;
    node [style=filled];
    "Idle" -> "Init" [ label = "Start" ];
    "Red" -> "Yellow" [ label = "CritReady" ];
    "Green" -> "Red" [ label = "CritFail" ];
    "Yellow" -> "Red" [ label = "CritFail" ];
    "Red" -> "Red" [ label = "CritFail" ];
    "Yellow" -> "Green" [ label = "OptReady" ];
    "Green" -> "Yellow" [ label = "OptFail" ];
    "Init" -> "Red" [ label = "ModuleReady" ];
    "Yellow" -> "Yellow" [ label = "OptFail" ];

    "Idle" [color="0.000 0.000 0.900"];
    "Init" [color="0.000 0.000 0.600"];
    "Red" [color="red"];
    "Yellow" [color="yellow"];
    "Green" [color="green"];
}
```