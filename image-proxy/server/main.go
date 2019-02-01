package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	pb "github.com/jun06t/grpc-sample/image-proxy/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterConverterServer(s, new(server))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) Convert(stream pb.Converter_ConvertServer) error {
	id, src, qa, err := receive(stream)
	if err != nil {
		return err
	}

	dst := fmt.Sprintf("%s.webp", id)

	cmd := exec.Command("cwebp", "-quiet", "-q", qa, "-o", dst, src)
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = send(stream, dst)
	if err != nil {
		return err
	}

	return nil
}

func receive(stream pb.Converter_ConvertServer) (string, string, string, error) {
	var (
		id   string
		qa   string
		name string
		file *os.File
	)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", "", "", err
		}
		if meta := resp.GetMeta(); meta != nil {
			name = fmt.Sprintf("%s.%s", meta.Id, meta.Type)
			qa = meta.Quality
			file, err = os.Create(name)
			if err != nil {
				return "", "", "", err
			}
			defer file.Close()
		}
		if chunk := resp.GetChunk(); chunk != nil {
			file.Write(chunk.Data)
		}
	}

	return id, name, qa, nil
}

func send(stream pb.Converter_ConvertServer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		data := &pb.ConvertResponse{
			Data:     buf,
			Position: int64(n),
		}
		stream.Send(data)
	}

	return nil
}
