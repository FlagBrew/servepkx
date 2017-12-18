package clientPKSM;

import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.Socket;

public class Client {

	private final String header = "PKSMOTA";
	private final int pk7size = 232;
	private final int maxPkxPerSocket = 30;
	
	private Socket socket;
	private int port;
	private String host;
	private ByteArrayOutputStream baos;
	
	private Client()
	{
		baos = new ByteArrayOutputStream();
	}

	public Client(String host)
	{
		this();
		this.host=host;
		this.port=9000;
	}
	
	public Client(String host, int port)
	{
		this();
		this.port = port;
		this.host = host;
	}
	
	public boolean queue(File file)
	{
		boolean ret = false;
		try
		{
			byte[] buffer = new byte[1024];
			InputStream in = new FileInputStream(file);
			int count = in.read(buffer);
			
			if (count == pk7size && baos.size() < 232*maxPkxPerSocket)
			{
				baos.write(buffer, 0, count);
				ret = true;
			}
			
			in.close();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();			
		}
		
		return ret;
	}

	public void send()
	{
		try
		{
			this.socket = new Socket(this.host, this.port);
			OutputStream out = socket.getOutputStream();
			out.write(header.getBytes(), 0, header.length());
			out.write(baos.toByteArray(), 0, baos.size());
			out.close();
			socket.close();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();	
		}
	}
}
