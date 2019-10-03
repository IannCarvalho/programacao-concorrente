package questao4;

import java.io.IOException;
import java.io.Serializable;
import java.util.Arrays;

@SuppressWarnings("serial")
public class MyHashMap implements Serializable{
	private Persistence banco;
	private Integer max;
	private Integer maxTime;
	private Integer[] hashMap;

	public MyHashMap(Integer max, Integer maxTime) {
		this.max = max;
		this.maxTime = maxTime;
		this.hashMap = new Integer[max];
		this.banco = new Persistence();
	}
	
	public Integer[] getHashMap() {
		return hashMap;
	}

	public Integer getHashKey(Integer num) {
		return num % max;
	}

	public void clear() {
		this.hashMap = new Integer[max];
	}

	public boolean containsKey(Integer key) throws ClassNotFoundException, IOException {
		Integer hashKey = getHashKey(key);
		Integer result = hashMap[hashKey];
		if (result == null)
			result = banco.recuperar()[hashKey];
		return (result == null);
	}

	public Integer get(Integer key) throws ClassNotFoundException, IOException {
		Integer hashKey = getHashKey(key);
		Integer value = hashMap[hashKey];
		if (value == null)
			value = banco.recuperar()[hashKey];
		return value;
	}

	// Só local
	public boolean isEmpty() throws ClassNotFoundException, IOException {
		boolean isEmpty = true;
		int i = 0;

		while (isEmpty && i < max) {
			isEmpty = (hashMap[i] == null);
			i++;
		}

		return isEmpty;
	}

	public Integer put(Integer key, Integer value) throws ClassNotFoundException, IOException {
		Integer hashKey = getHashKey(key);
		Integer previousValue = hashMap[hashKey];
		if (previousValue == null)
			previousValue = banco.recuperar()[hashKey];
		hashMap[hashKey] = value;
		return previousValue;
	}

	public Integer remove(Integer key, Integer value) throws ClassNotFoundException, IOException {
		Integer hashKey = getHashKey(key);
		Integer previousValue = hashMap[hashKey];
		if (previousValue == null)
			previousValue = banco.recuperar()[hashKey];
		hashMap[hashKey] = null;
		return previousValue;
	}

	// Só para o local
	public Integer size() throws ClassNotFoundException, IOException {
		int counter = 0;
		
		for (int i = 0; i < max; i++) {
			if (hashMap[i] != null)
				counter++;
		}

		return counter;
	}
	
	public String toString() {
		return Arrays.toString(hashMap);
	}
}
