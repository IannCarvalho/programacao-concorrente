# Questão 4

Considere um programa em java que usa um banco de dados chave-valor que implementa a interface [Map](https://docs.oracle.com/javase/8/docs/api/java/util/Map.html). Operações que alteram o estado desse banco (p.ex put e remove) são bastante lentas; uma vez que os dados modificações precisam ser enviados para o disco de maneira síncrona. Operações de leitura são mais rápidas (entretanto, mais lentas que em uma implementação de Map, digamos HashMap, que mantém os dados em memória). Você deve implementar um cache como um Decorator --- em resumo, uma nova implementação de Map que repassa algumas chamadas para o banco de dados Map) ---- para acelerar o desempenho. Essa sua implementação precisa ser thread-safe e não relaxar a consistência da estrutura de dados. Múltiplas threads podem acessar o eu CacheMap.

O construtor do seu Map deve receber três parâmetros:
 1. uma referência para banco de dados Map;
 2. um número inteiro que indica a quantidade máxima de pares chave-valor que podem ser mantidos em memória por sua implementação;
 3. um inteiro que indica o tempo máximo, em segundos, que o banco pode ficar sem ser atualizado.
 
 Da interface de Map, implemente os seguintes métodos:

* clear
* containsKey
* get
* isEmpty
* put
* remove(key, object)
* size

## Decisões de implementação

1. Foi decidido que ao dar timeout o cache não seria limpo, pois caso seja definido um número máximo de pares grande (100, por exemplo), ao existir 50 pares chave-valor no cache, não é interessante eliminar todos essess pares pos tornariam os gets mais custosos.

2. O método `size()` retorna o tamanho do cache unido ao tamanho do banco.

3. Para que haja uma forma de interromper a execução da Thread de atualização, foi incluido o atributo `shouldUpdate`, inicialmente `true` e que, ao se tornar `false`, quebra o while da Thread.

## Como executar

Por haver sido setado um tempo relativamente longo de timeout (3 segundos), os testes demoram relativamente muito a serem executados. Esse timeout pode ser alterado na classe CachedMapTest.

```shell script
mvn clean test
```