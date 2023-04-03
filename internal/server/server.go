package server

type server interface {
	ListenAndServe() error
}

func StartServer(server server) error {
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
