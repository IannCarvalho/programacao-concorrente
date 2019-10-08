package q4;

import java.util.Calendar;
import java.util.HashMap;
import java.util.Map;
import java.util.HashSet;
import java.util.Set;

public class MyHashMap {
	
	private int size;
	
	private int maxSize;
	
	private int currentSize;
	
	private Calendar start;
	
	private int timeout;
	
	private Set<Integer> keys;
	
	private Map<Integer, Integer> cache;
	
	private DatabaseAccess dbAccess;
	
	public MyHashMap(int maxSize, int timeout) {
		this.size = 0;
		this.currentSize = 0;
		this.maxSize = maxSize;
		this.timeout = timeout;
		this.cache = new HashMap<Integer, Integer>();
		this.dbAccess = new DatabaseAccess("");
		this.start = Calendar.getInstance();
		this.keys = new HashSet<Integer>();
	}

	public int size() {
		return size;
	}

	public boolean isEmpty() {
		return size == 0;
	}

	public boolean containsKey(Integer key) {
		return this.cache.containsKey(key) || this.keys.contains(key);
	}

	public Integer get(Integer key) {
		Integer value = null;
		if (cache.containsKey(key)) {
			value = cache.get(key);
		} else if (this.containsKey(key)) {
			value = dbAccess.getDbValue(key);
		}
		return value;
	}

	public synchronized Integer put(Integer key, Integer value) {
		Integer oldValue = null;
		this.currentSize++;
		if (!this.containsKey(key)) {
			this.size++;
		} else if (!this.cache.containsKey(key)) {
			oldValue = dbAccess.getDbValue(key);
		} else {
			oldValue = this.cache.get(key);
		}
		this.cache.put(key, value);
		return oldValue;
	}

	public synchronized Integer remove(Integer key) {
		Integer valueToRemove = null;
		this.currentSize++;
		if (this.containsKey(key)) {
			this.size--;
			valueToRemove = this.get(key);
		}
		this.verifyAndTriggerUpdate();
		return valueToRemove;
	}

	public synchronized void clear() {
		this.clearCache();
		this.dbAccess.reset();
	}
	
	private synchronized void clearCache() {
		this.cache = new HashMap<Integer, Integer>();
	}
	
	private synchronized void verifyAndTriggerUpdate() {
		int nowMilliseconds = Calendar.getInstance().get(Calendar.MILLISECOND);
		int startMilliseconds = this.start.get(Calendar.MILLISECOND); 
		if (this.currentSize >= this.maxSize ||
				 nowMilliseconds >= startMilliseconds + (timeout * 1000))
			this.dbAccess.updateDatabase(cache);
	}

}
