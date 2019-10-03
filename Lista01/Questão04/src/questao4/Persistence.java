package questao4;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;

public class Persistence {
	public void gravar(Integer[] hm) throws FileNotFoundException, IOException{
		ObjectOutputStream saida = new ObjectOutputStream(new FileOutputStream("MyHashMap.obj"));
		saida.writeObject(hm);
		saida.close();
	}
	
	public Integer[] recuperar() throws ClassNotFoundException, IOException{
		ObjectInputStream entrada = new ObjectInputStream(new FileInputStream("MyHashMap.obj"));
		Integer[] hm = (Integer[]) entrada.readObject();
		entrada.close();
		return hm;
	}
}
