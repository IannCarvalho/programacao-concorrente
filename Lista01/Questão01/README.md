# Questão 1
Em sala, discutimos uma implementação de um **lock** justo, em clang. Esse **lock** usava uma fila para manter identificadores de **pthreads** em espera para executar a região crítica. Uma vez liberado o **lock**, a thread que entrou primeiro na fila deve ser escolhida para executar. Faça uma implementação em java com as mesmas características, seguindo a interface listada abaixo. Sua implementação não pode usar os métodos **wait**, **notify** e **notifyAll** da classe Object. Use **ArrayList** para implementar a fila. Você não pode usar **synchronized** na declaração de nenhum método criado por você. Você não pode usar nenhum objeto do pacote **java.util.concurrent** exceto **java.util.concurrent.locks.LockSupport.park()** para bloquear a execução da Thread corrente e **java.util.concurrent.locks.LockSupport.unpark(Thread thread)** para desbloquear. Adicionalmente, você pode usar **java.util.concurrent.atomic.AtomicInteger** para implementar o equivalente à instrução testAndSet caso topo não usar **synchronized** em nenhum ponto do código (nem em um bloco interno).

## Decisões de implementação

1. Foi decidido usar uma classe que era responsável pelas manipulações na lista e de lock, chamada "Fifo"

2. Todas as threads passam por um processo de lock depois de serem criadas, no qual são adicionadas na fila de forma sicronizada e depois verifica-se se a thread que está tentando rodar é a que tá no começo da fila, caso seja, ela roda o código de run, caso não seja ela é colocada em *park*

3. Depois de rodar o código de run, a thread que finalizou seu processo, fará o unlock, o qual remove ela da fila e depois tira o próximo elemento do *park* através do *unpark*

## Como executar

```
mvn clean compile exec:java
```
