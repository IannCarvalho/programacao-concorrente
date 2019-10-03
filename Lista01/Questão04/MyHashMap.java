package questao4;

import java.util.Arrays;
import java.util.LinkedList;

public class MyHashMap {
	// private algumaCoisa banco;
	private Integer max;
	//private Integer maxTime;
	private LinkedList<Integer>[] hashMap;

	public MyHashMap(Integer max, Integer maxTime) {
		this.max = max;
		//this.maxTime = maxTime;
		clear();
	}

	public Integer getHashKey(Integer num) {
		return num % max;
	}

	@SuppressWarnings("unchecked")
	public void clear() {
		this.hashMap = new LinkedList[max];
		for (int i = 0; i < this.max; i++)
			hashMap[i] = new LinkedList<Integer>();
	}

	public boolean containsKey(Integer key) {
		Integer hashKey = getHashKey(key);
		return hashMap[hashKey].isEmpty();
	}

	public LinkedList<Integer> get(Integer key) {
		Integer hashKey = getHashKey(key);
		return hashMap[hashKey];
	}

	public boolean isEmpty() {
		boolean isEmpty = true;
		int i = 0;

		while (isEmpty && i < max) {
			isEmpty = hashMap[i].isEmpty();
			i++;
		}

		return isEmpty;
	}

	public LinkedList<Integer> put(Integer key, Integer value) {
		LinkedList<Integer> previousValue = get(key);
		previousValue.add(value);
		return previousValue;
	}

	public LinkedList<Integer> remove(Integer key, Integer value) {
		LinkedList<Integer> previousValue = get(key);
		previousValue.remove(value);
		return previousValue;
	}

	public Integer size() {
		int counter = 0;

		for (int i = 0; i < max; i++) {
			if (!hashMap[i].isEmpty())
				counter++;
		}

		return counter;
	}
	
	public String toString() {
		return Arrays.toString(this.hashMap);
	}
}
