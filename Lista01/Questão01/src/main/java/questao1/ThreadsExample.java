package questao1;

public class ThreadsExample {
	public static void main(String[] args) {
		Fifo fifo = new Fifo();

		// As threads podem não rodar na ordem que são criadas, mas rodam na ordem que
		// estão sendo adicionadas no array
		for (int i = 0; i < 100; i++) {
			fifo.runThread(Integer.toString(i + 1));
		}
	}
}