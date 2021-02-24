package main

import (
	"context"
	"flag"
	"fmt"
	dgo "github.com/dgraph-io/dgo/v2"
	api "github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
	"log"
)


type Person struct{
	Uid		string	`json:"uid,omitempty"`
	Name		string	`json:"name,omitempty"`
	From		string	`json:"from,omitempty"`
	NameOFcn	string	`json:"nameOFcn,omitempty"`
	NameOFjp	string	`json:"nameOFjp,omitempty"`
	NameOFen	string	`json:"nameOFen,omitempty"`
	Age		int	`json:"age,omitempty"`
	Friend		[]Person `json:"friend,omitempty"`
}

var (
	dgraph = flag.String("d", "10.34.11.90:9080", "Dgraph server address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	//p1 := Person{
	//	Name: "wangha",
	//	Age: 17,
	//	From: "China",
	//	NameOFen: "wangha",
	//	NameOFcn: "王哈",
	//	NameOFjp: "王ハ",
	//}
	//p2 := Person{
	//	Name: "chenchao",
	//	Age: 22,
	//	From: "China",
	//	NameOFen: "ChaoChen",
	//	NameOFcn: "陈超",
	//}
	//p3 := Person{
	//	Name: "xhe",
	//	Age: 18,
	//	From: "Japan",
	//	NameOFen: "wanghe",
	//	NameOFcn: "x鹤",
	//}
	//p4 := Person{
	//	Name: "changyang",
	//	Age: 19,
	//	From: "England",
	//	NameOFcn: "常飏",
	//}
	//p5 := Person{
	//	Name: "yetao",
	//	Age: 18,
	//	From: "Russian",
	//	NameOFen: "TaoYe",
	//	NameOFcn: "叶掏",
	//}

	//op := &api.Operation{}
	//op.Schema = `
	//	uid:  string .
	//	name: string .
	//	age: int .
	//	from: string .
	//	nameOFcn: string @index(term) .
	//	nameOFjp: string @index(term) .
	//	nameOFen: string @index(term) .
	//`
	//
	ctx := context.Background()
	op := &api.Operation{
		Schema: `name: string @index(hash) .`,
	}
	if err := dg.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}
	//
	//mu := &api.Mutation{
	//	CommitNow: true,
	//}

	//var p = [5]Person{p1,p2,p3,p4,p5}
	//
	//for _,x := range p {
	//	pb, err := json.Marshal(x)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	mu.SetJson = pb
	//	_,err = dg.NewTxn().Mutate(ctx, mu)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}
	txn := dg.NewTxn()
	defer txn.Discard(ctx)
	q := `query all($a: string) {
    all(func: eq(name, $a)) {
       	uid
		name
		age
		from
		nameOFcn
		nameOFjp
		nameOFen
    }
  }`

	res, err := txn.QueryWithVars(ctx, q, map[string]string{"$a": "changyang"})
	fmt.Printf("%s\n", res.Json)

	req := &api.Request{
		Query: q,
		Vars: map[string]string{"$a": "xhe"},
	}
	res, err = txn.Do(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", res.Json)

	//query := `
	//  query {
	//	  user as var(func: eq(email, "wrong_email@dgraph.io"))
	//  }`
	//mu := &api.Mutation{
	//	SetNquads: []byte(`uid(user) <email> "correct_email@dgraph.io" .`),
	//}
	//req = &api.Request{
	//	Query: query,
	//	Mutations: []*api.Mutation{mu},
	//	CommitNow:true,
	//}
	//
	//// Update email only if matching uid found.
	//if _, err := dg.NewTxn().Do(ctx, req); err != nil {
	//	log.Fatal(err)
	//}

}
