## This package defines a rudimentary init FSM

### State Transition Table
| Input\Current State | Idle | Red    | Yellow | Green  |
| ------------------- | ---- | ------ | ------ | ------ |
| StartM              | Red  | ...... | ...... | ...... |
| VitalReady          | .... | Yellow | Yellow | ...... |
| VitalFail           | .... | Red    | Red    | Red    |
| EssReady            | .... | ...... | Green  | Green  |
| EssFail             | .... | ...... | Yellow | Yellow |
| Shutdown            | Stop | Stop   | Stop   | Stop   |

### Digraph
```
digraph fsm {
    ratio = fill;
    node [style=filled];
    
    subgraph cluster_0 {
        style=filled;
        color=lightgrey;
        node [style=filled,color=white];
        Red -> Yellow [label = "VitalReady"];
        Yellow -> Red [label = "VitalFail"];
        Yellow -> Green [label = "EssReady"];
        Green -> Yellow [label = "EssFail"]
        Green -> Red [label = "VitalFail"]
        
        Red -> Red [label = "VitalFail"]
        Yellow -> Yellow [label = "EssFail"]
        Yellow -> Yellow [label = "VitalReady"]
        Green -> Green [ label = "EssReady" ];
        label = "";
    }
    
    "Red" [color="red"];
    "Yellow" [color="yellow"];
    "Green" [color="green"];
    "Idle" [color="0.000 0.000 0.900"];
    "Stop" [color="0.000 0.000 0.600"];
    
    "Idle" -> "Red" [label = "StartM"]
    "Green" -> "Stop" [label = "Shutdown"]
    "Idle" -> "Stop" [label = "Shutdown"]
    "Yellow" -> "Stop" [label = "Shutdown"]
    "Red" -> "Stop" [label = "Shutdown"]
    
}
```