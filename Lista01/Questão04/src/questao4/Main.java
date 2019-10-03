package questao4;

import java.io.IOException;
import java.util.Arrays;

public class Main implements Runnable {
	public static void main(String args[]) throws ClassNotFoundException, IOException {
		for (int i = 0; i < 5; i++) {
			Thread t = new Thread(new Main(), Integer.toString(i));
			t.start();
		}
		MyHashMap myHashMap = new MyHashMap(10, 10);
		Persistence persistence = new Persistence();
		persistence.gravar(myHashMap.getHashMap());
		
		myHashMap.put(1, 1);
		myHashMap.put(1, 2);
		myHashMap.put(1, 1);
		myHashMap.put(1, 1);
		myHashMap.put(3, 1);
		myHashMap.put(1, 1);
		
		System.out.println(myHashMap);
		persistence.gravar(myHashMap.getHashMap());
		myHashMap.clear();
		Integer[] banco = persistence.recuperar();
		System.out.println(myHashMap);
		System.out.println(Arrays.toString(banco));
	}

	@Override
	public void run() {
		Thread current = Thread.currentThread();

		if(current.getName().equals("0")) {
		}
		if(current.getName().equals("1")) {}
		if(current.getName().equals("2")) {}
	}
}
