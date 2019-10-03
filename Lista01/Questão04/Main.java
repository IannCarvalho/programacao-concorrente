package questao4;

public class Main {
	public static void main(String args[]) {
		MyHashMap myHashMap = new MyHashMap(10, 10);
		myHashMap.put(1, 1);
		myHashMap.put(1, 2);
		myHashMap.put(1, 1);
		myHashMap.put(1, 1);
		myHashMap.put(3, 1);
		myHashMap.put(1, 1);
	}
}
