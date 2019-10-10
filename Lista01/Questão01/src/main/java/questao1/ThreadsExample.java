package questao1;

public class ThreadsExample{
    public static void main(String[] args) {
    	Fifo fifo = new Fifo();
    	
    	for(int i = 0; i < 100; i++) {    		
    		fifo.runThread(Integer.toString(i+1));
    	}
    }
}