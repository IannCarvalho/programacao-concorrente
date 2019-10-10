# Questão 2

CountDownLatch é um sincronizador disponível na sdk Java (https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/CountDownLatch.html). Faça uma nova implementação da mesma ideia usando somente variáveis condicionais (**wait**, **notify** e **notifyAll**) e a construção **synchronized**. Se precisar usar alguma coleção, use uma LinkedList ou ArrayList. Só é necessário implementar a API abaixo:
* void await()
* void countDown()

## Decisões de implementação

1. Foi decidido que usar uma classe que era responsável só pelo latch que servia para fazer o countDown e o await e toda a logística de notificar e colocar as threads para dormir

2. Worker foi uma classe que representava todas as threads.

3. Todos os Workers possuiam o mesmo latch.

## Como executar

```
mvn clean compile exec:java
```
