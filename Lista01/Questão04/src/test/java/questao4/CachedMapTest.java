package questao4;

import static org.junit.Assert.*;
import org.junit.Before;
import org.junit.Test;

import java.util.HashMap;
import java.util.Map;

public class CachedMapTest {

    private static final int TIMEOUT = 3;

    private static final int PAST_TIMEOUT = TIMEOUT + 1;

    private static final int MAX_PAIRS = 5;

    private CachedMap cache;

    private Map<Integer, Integer> disco;

    @Before
    public void setUp() {
        disco = new HashMap<Integer, Integer>();
        cache = new CachedMap(MAX_PAIRS, TIMEOUT, disco);
    }

    /**
     * Testa o fluxo de adicionar um par chave-valor, verificar
     * se o par está no cache e não está no banco em disco e, após
     * o timeout, verifica se o par estará em disco e no cache.
     */
    @Test
    public void testAddingOne() throws InterruptedException {
        cache.put(1, 1);
        assertEquals(0, disco.size());
        assertEquals(1, cache.size());

        Thread.sleep(PAST_TIMEOUT * 1000);
        cache.deactivate();

        assertEquals(1, disco.size());
        assertEquals(1, cache.size());
    }

    /**
     * Testa o fluxo de adicionar cinco pares chave-valor, verificar
     * se os pares estão no cache e não estão no banco em disco e, após
     * o timeout, verifica se estarão em disco e no cache.
     */
    @Test
    public void testAddingBelowLimit() throws InterruptedException {
        for (int i = 0; i < MAX_PAIRS; i++) {
            cache.put(i, 1);
        }

        assertEquals(0, disco.size());
        assertEquals(5, cache.size());

        Thread.sleep(PAST_TIMEOUT * 1000);
        cache.deactivate();

        assertEquals(5, disco.size());
        assertFalse(cache.isCacheEmpty());
    }

    /**
     * Testa o fluxo de adicionar seis pares chave-valor, verificar
     * se os pares estão no banco em disco e não estão no cache e, após
     * o timeout, verifica se ainda estarão em disco e não no cache.
     */
    @Test
    public void testAddingLimit() throws InterruptedException {
        for (int i = 0; i < MAX_PAIRS + 1; i++) {
            cache.put(i, 1);
        }

        assertEquals(6, disco.size());
        assertTrue(cache.isCacheEmpty());

        Thread.sleep(PAST_TIMEOUT * 1000);
        cache.deactivate();

        assertEquals(6, disco.size());
        assertTrue(cache.isCacheEmpty());
    }

    /**
     * Testa o fluxo de adicionar um par chave-valor e, antes do timeout,
     * remover o par inserido. Além disso, espera o timeout para certificar-se
     * que o par não foi enviado ao banco depois de ser removido do cache.
     */
    @Test
    public void testAddOneAndRemoveIt() throws InterruptedException {
        cache.put(1, 1);
        assertEquals(0, disco.size());
        assertEquals(1, cache.size());

        cache.remove(1);

        assertEquals(0, disco.size());
        assertEquals(0, cache.size());

        Thread.sleep(PAST_TIMEOUT * 1000);
        cache.deactivate();

        assertEquals(0, disco.size());
        assertEquals(0, cache.size());
    }

    /**
     * Testa o fluxo de adicionar cinco pares chave-valor e, antes do timeout,
     * remover os pares inseridos. Além disso, espera o timeout para certificar-se
     * que o pares não foram enviados ao banco depois de serem removido do cache.
     */
    @Test
    public void testAddBelowLimitWaitTimeoutAndRemove() throws InterruptedException {
        for (int i = 0; i < MAX_PAIRS; i++) {
            cache.put(i, 1);
        }

        assertEquals(0, disco.size());
        assertEquals(5, cache.size());

        for (int i = 0; i < MAX_PAIRS; i++) {
            cache.remove(i);
        }

        assertEquals(0, disco.size());
        assertEquals(0, cache.size());

        Thread.sleep(PAST_TIMEOUT * 1000);
        cache.deactivate();

        assertTrue(disco.isEmpty());
        assertTrue(cache.isCacheEmpty());
    }

    /**
     * Testa o fluxo de adicionar seis pares chave-valor, executando um
     * flush para o banco, e, após o timeout, remover um dos pares
     * inseridos, certificando-se que o par removido não foi adicionado
     * ao cache e o par não está mais no disco
     */
    @Test
    public void testAddLimitAndRemoveFromDisk() throws InterruptedException {
        for (int i = 0; i < MAX_PAIRS + 1; i++) {
            cache.put(i, 1);
        }

        assertEquals(6, disco.size());
        assertTrue(cache.isCacheEmpty());

        Thread.sleep(PAST_TIMEOUT * 1000);

        assertEquals(6, disco.size());
        assertTrue(cache.isCacheEmpty());

        cache.remove(1);
        assertEquals(6, disco.size());
        assertTrue(cache.isCacheEmpty());

        Thread.sleep(PAST_TIMEOUT * 1000);
        cache.deactivate();

        assertEquals(5, disco.size());
        assertTrue(cache.isCacheEmpty());
    }
}
