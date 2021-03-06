package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"

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
	buf := pool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		pool.Put(buf)
	}()

	m, err := receive(stream, buf)
	if err != nil {
		return err
	}

	// NOTE: generate uuid when each call
	uuid := "uuid-"
	src := uuid + "src"
	defer os.Remove(src)

	err = writeOrg(src, buf)
	if err != nil {
		return err
	}

	dst := uuid + "dst"
	defer os.Remove(dst)

	cmd := exec.Command("cwebp", "-quiet", "-mt", "-q", m.qa, "-o", dst, src)
	err = cmd.Run()
	if err != nil {
		return err
	}

	file, err := os.Open(dst)
	if err != nil {
		return err
	}
	defer file.Close()
	err = send(stream, file)
	if err != nil {
		return err
	}

	return nil
}

var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1024*64))
	},
}

type meta struct {
	qa string
}

func receive(stream pb.Converter_ConvertServer, w io.Writer) (meta, error) {
	m := meta{}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return m, err
		}

		if mt := resp.GetMeta(); mt != nil {
			m.qa = mt.Quality
		}
		if chunk := resp.GetChunk(); chunk != nil {
			_, err := w.Write(chunk.Data)
			if err != nil {
				return m, err
			}
		}
	}

	return m, nil
}

func writeOrg(name string, src io.Reader) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, src)
	if err != nil {
		return err
	}

	return nil
}

func send(stream pb.Converter_ConvertServer, r io.Reader) error {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
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
