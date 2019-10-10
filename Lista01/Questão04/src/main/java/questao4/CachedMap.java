package questao4;

import java.util.*;
import java.util.concurrent.TimeUnit;

public class CachedMap {

    private int maxSize;

    private volatile Set<Integer> keys;

    private volatile Set<Integer> keysToRemove;

    private volatile Map<Integer, Integer> cache;

    private volatile Map<Integer, Integer> dbAccess;

    private volatile boolean shouldUpdate;

    public CachedMap(int maxSize, int timeout, Map<Integer, Integer> db) {
        this.maxSize = maxSize;
        this.cache = new HashMap<Integer, Integer>();
        this.dbAccess = db;
        this.keysToRemove = new HashSet<Integer>();
        this.keys = new HashSet<Integer>();
        this.keys.addAll(db.keySet());
        this.shouldUpdate = true;
        Thread t = new Thread(new VerifyAndUpdate(this, timeout), "THREAD_UPDATE");
        t.start();
    }

    public int size() {
        return keys.size();
    }

    public boolean isCacheEmpty() {
        return this.cache.isEmpty();
    }

    public boolean isEmpty() {
        return keys.size() == 0;
    }

    public synchronized void deactivate() {
        this.shouldUpdate = false;
    }

    public boolean containsKey(Integer key) {
        return this.keys.contains(key);
    }

    public Integer get(Integer key) {
        Integer value = null;
        if (cache.containsKey(key)) {
            value = cache.get(key);
        } else if (this.containsKey(key)) {
            value = dbAccess.get(key);
            this.put(key, value);
        }
        return value;
    }

    public synchronized Integer put(Integer key, Integer value) {
        Integer oldValue = null;
        if (!this.containsKey(key)) {
            this.keys.add(key);
        } else if (this.cache.containsKey(key)) {
            oldValue = this.cache.get(key);
        } else {
            oldValue = dbAccess.get(key);
        }

        this.cache.put(key, value);

        if (this.keysToRemove.contains(key)) {
            this.keysToRemove.remove(key);
        }

        this.verifyAndTrigger();
        return oldValue;
    }

    public synchronized Integer remove(Integer key) {
        Integer valueToRemove = null;
        if (this.containsKey(key)) {
            this.keys.remove(key);
            valueToRemove = this.get(key);
            cache.remove(key);
            keysToRemove.add(key);
        }
        this.verifyAndTrigger();
        return valueToRemove;
    }

    public synchronized void clear() {
        this.clearCache();
        this.dbAccess.clear();
    }

    private synchronized void clearCache() {
        this.cache.clear();
    }

    private synchronized void flush() {
        if (this.shouldUpdate) {
            this.dbAccess.putAll(cache);
            for (Integer key: this.keysToRemove) {
                dbAccess.remove(key);
            }
            keysToRemove.clear();
        }
    }

    private synchronized void verifyAndTrigger() {
        if (cache.size() > maxSize) {
            this.flush();
            this.clearCache();
        }
    }

    private class VerifyAndUpdate implements Runnable {

        private int timeout;

        private CachedMap cache;

        VerifyAndUpdate(CachedMap cache, int timeout) {
            this.timeout = timeout;
            this.cache = cache;
        }

        public void run() {
            while (cache.shouldUpdate) {
                try {
                    TimeUnit.SECONDS.sleep(timeout);
                    cache.flush();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }
    }

}
