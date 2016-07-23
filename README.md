By default, only one job and work when run script, otherwise, it's possible more fast increasing jobs or works. 
Access file `main.go` and edit follow lines:
```go
const numberOfJobs = 1
const numberOfWorks = 1
```

Run follow command:

```shell
$ time go run main.go
{Shanghai International Circuit}
{Circuit de Spa-Francorchamps}
{Circuit Gilles Villeneuve}
{Sepang International Circuit}
{Silverstone Circuit}
{Circuit de Monaco}
{Marina Bay Street Circuit}
{Hungaroring}
{Autódromo José Carlos Pace}
{Hockenheimring}
{Circuit of the Americas}
{Bahrain International Circuit}
{Red Bull Ring}
{Circuit de Catalunya}
{Suzuka Circuit}
{Autódromo Hermanos Rodríguez}
{Sochi International Street Circuit}
{Albert Park Grand Prix Circuit}
{Autodromo Nazionale di Monza}
{Baku City Circuit}

go run main.go  1.05s user 0.13s system 61% cpu 1.912 total
```

