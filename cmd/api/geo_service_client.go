package main

import (
	"context"
	"fmt"
	"geotask_pprof/geo"
	"log"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("localhost", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := geo.NewGeoServiceClient(conn)

	searchRequest := &geo.SearchRequest{}
	searchResponse, err := client.AddressSearch(context.Background(), searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AddressSearch Response:", searchResponse)

	geoCodeRequest := &geo.GeoCodeRequest{
		Lat: "Latitude",
		Lng: "Longitude",
	}
	geoCodeResponse, err := client.GeoCode(context.Background(), geoCodeRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GeoCode Response:", geoCodeResponse)
}
