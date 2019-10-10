### Questão 3 - Threads, processos leves. Escreve dois programas em clang. O primeiro, cria N threads e em seguida, executa join para cada um delas. Cada thread deve simplesmente dormir por um tempo (alguns segundos), e em seguida executar thread_exit(0). Faça um segundo programa, equivalente ao anterior, que cria N processos, que dormem por um tempo determinado, e os esperar terminar. Avalie as diferenças de desempenho, tanto em tempo decorrido para executar as operações importantes), quanto no consumo de memória.


#### Solução para cálculo de tempo e custo

* Como dito no enunciado da questão, os programas geram N threads e processos, nós escolhemos 1000.
* Para que não houvessem gastos de memória com IO, optou-se por gerar dois programas para cada problema. Um deles responsável por calcular o tempo e fazer a persistência desse cálculo, para uma posterior análise descritiva comparativa. Outro sendo uma versão simplificada do código anterior que não lida com IO e nem cálculo de tempo, para mitigar o consumo de memória periférico.
* Os códigos aqui se encontram para as soluções de processos para [memória](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ProcessMemory.C) e para [tempo](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ProcessTime.C)
* Os códigos aqui se encontram para as soluções de threads para [memória](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ThreadMemory.C) e para [tempo](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ThreadTime.C)

#### Por que as threads gastam menos tempo que processos?
![Tempo1](https://imagizer.imageshack.com/img921/9671/Pr0JpP.png)

Imagem 01 - Gráfico comparando o tempo de execução das threads e dos processos

#### Por que as threads gastam menos recursos que processos?

![Consumo1](https://imagizer.imageshack.com/img921/463/yR4soa.png)

Imagem 02 - O consumo de memória de execução das threads

![Consumo2](https://imagizer.imageshack.com/img924/5390/Z4coBN.png)

Imagem 03 - O consumo de memória de execução dos processos
