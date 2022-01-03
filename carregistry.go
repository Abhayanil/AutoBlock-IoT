package main


import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


type VehicleAccidentRegistery struct {
}

type vehicle struct {
	VehNo   int `json:"vehno"`
	gyro_xaxis  float32 `json:"gyro_xaxis"`
	gyro_yaxis  float32 `json:"gyro_yaxis"`
	gyro_zaxis  float32 `json:"gyro_zaxis"`
	gps_lat float32 `json:"gps_lat"`
	gps_long  float32 `json:"gps_long"`
}


func (s *VehicleAccidentRegistery) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}


func (s *VehicleAccidentRegistery) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	
	function, args := stub.GetFunctionAndParameters()
	
	if function == "vehicledetails" {

		return s.vehicledetails(stub, args)

	} else if function == "newvehicle" {

		return s.newvehicle(stub, args)

	} 

	return shim.Error("Invalid function name.")
}

func (s *VehicleAccidentRegistery) vehicledetails(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	vehicleAsBytes, _ := stub.GetState(args[0])

	return shim.Success(vehicleAsBytes)
}
 

func (s *VehicleAccidentRegistery) newvehicle(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var vehicle = Vehicle{VehNo: args[1], gyro_xaxis: args[2], gyro_yaxis: args[3], gyro_zaxis: args[4], gps_lat: args[5], gps_long: args[6]}

	vehicleAsBytes, _ := json.Marshal(vehicle)

	stub.PutState(args[0], vehicleAsBytes)

	return shim.Success(nil)
}


func main() {

	err := shim.Start(new(VehicleAccidentRegistery))
	if err != nil {
		fmt.Printf("Error creating new chaincode: %s", err)
	}
}