### Questão 3 - Threads, processos leves. Escreve dois programas em clang. O primeiro, cria N threads e em seguida, executa join para cada um delas. Cada thread deve simplesmente dormir por um tempo (alguns segundos), e em seguida executar thread_exit(0). Faça um segundo programa, equivalente ao anterior, que cria N processos, que dormem por um tempo determinado, e os esperar terminar. Avalie as diferenças de desempenho, tanto em tempo decorrido para executar as operações importantes), quanto no consumo de memória.

#### Introdução teórica

##### Como funcionam os processos?

Sabendo-se que um programa precisa de memória e outros recursos do sistema operacional para ser executado, um processo é o que chamamos de programa carregado na memória, juntamente com todos os recursos necessários para operar. O sistema operacional lida com a tarefa de gerenciar os recursos necessários para transformar seu programa em um processo em execução.

Em um computador, podem haver várias instâncias de um único programa, e cada instância desse programa em execução é um processo. Cada processo possui um espaço de endereço de memória separado, o que significa que um processo é executado de forma independente e é isolado de outros processos, ou seja, não é possível acessar diretamente os dados compartilhados em outros processos. Essa independência de processos é valiosa porque o sistema operacional faz o possível para isolar os processos, para que um problema em um processo não corrompa ou cause estragos em outro processo.

##### Como funcionam os threads?

Uma thread é a unidade de execução dentro de um processo. Quando um processo é iniciado, ele recebe memória e recursos e cada thread no processo compartilha essa memória e recursos.

Existem dois tipos de processos com relação a threads:
* os processos de thread única, que contém apenas thread, onde há apenas uma coisa acontecendo por vez.
* os processos multithread, que contem mais de uma thread, são aqueles processos realizam várias operações ao mesmo tempo.

Às vezes, as threads são chamados de processos leves porque têm sua própria pilha, mas podem acessar dados compartilhados. Como os threads compartilham o mesmo espaço de endereço, o custo operacional da comunicação entre os threads é baixo, o que é uma vantagem. A desvantagem é que um problema com um thread em um processo certamente afetará outros threads e a viabilidade do próprio processo.

#### Metodologia para solução para cálculo de tempo e custo

* Como dito no enunciado da questão, os programas geram N threads e processos, nós escolhemos 1000.
* Para que não houvessem gastos de memória com IO, optou-se por gerar dois programas para cada problema. Um deles responsável por calcular o tempo e fazer a persistência desse cálculo, para uma posterior análise descritiva comparativa. Outro sendo uma versão simplificada do código anterior que não lida com IO e nem cálculo de tempo, para mitigar o consumo de memória periférico.
* Os códigos aqui se encontram para as soluções de processos para [memória](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ProcessMemory.C) e para [tempo](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ProcessTime.C)
* Os códigos aqui se encontram para as soluções de threads para [memória](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ThreadMemory.C) e para [tempo](https://github.com/IannCarvalho/programacao-concorrente/blob/master/Lista01/Quest%C3%A3o03/ThreadTime.C)

#### Resultados

##### Tempo usado por threads e processos
![Tempo1](https://imagizer.imageshack.com/img921/9671/Pr0JpP.png)

Imagem 01 - Gráfico comparando o tempo de execução das threads e dos processos

No gráfico acima, podemos notar que o tempo de execução dos processos vai aumentando bem mais do que os das threads.

##### Memória usada por threads e processos

![Consumo1](https://imagizer.imageshack.com/img921/463/yR4soa.png)

Imagem 02 - O consumo de memória de execução das threads

![Consumo2](https://imagizer.imageshack.com/img924/5390/Z4coBN.png)

Imagem 03 - O consumo de memória de execução dos processos

Através das imagens acima percebemos que os proccessos usam mais memória do que a threads

##### Explicando os resultados

Como dito na secção de introdução teórica, fica evidente que o sistema operacional precisará de mais tempo e memória, tendo em vista que, diferente das threads, os processos necessitam de uma maior exclusividade dos recursos. Dessa forma, chegamos a conclusão de que os processos são mais custosos que as threads, apesar de serem mais seguros.
