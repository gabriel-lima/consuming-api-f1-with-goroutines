# Why?
More and more time I'm come studying the language Go. And, concurrency programming is a large topic anywhere.
Therefore, I created this small project for practice my studies.

# How Works?
The objective is consuming a API REST, in this case, I choice this API http://ergast.com/mrd/, because I'm a passionate for Formula 1. 
I need to get information of circuits of Formula 1 this year. 
Reading the API, I noted that she not provide a method to get all circuits in only one requisition.
I code a worker pools that request 20 circuits simultaneously, using goroutines and channels of language Go.

# Get Started
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

