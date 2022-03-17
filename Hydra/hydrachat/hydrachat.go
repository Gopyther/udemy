package hydrachat

var logger = hlogger.GetInstance()

//Start hydra chat
func Run(conneciton string) error {
	l, err := net.Listen("tcp", connection)
	if err != nil {
		logger.Println("Error connection to chat client", err)
		return err
	}
	r := CreateRoom("HydraChat")
}

go func() {
	//Handel SIGNT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<- ch

	l.close()
	fmt.Println("Closing tcp connection")
	close(r.Quit)
	if r.CLCount() > 0 {
		<- r.Msgch
	}
	os.Exit(0)
}()

for {
	conn, err := l.Accept()
	if err != nil {
		logger.Println("Error accepting connection from chat client", err)
		break
	}
	go handelConnection(r,conn)
}

retrun err