package apis

import(
	"time"
	"fmt"
	"net/http"
)

func CheckChallan(w http.ResponseWriter, r *http.Request){
	now := time.Now()
	vehicleNumber := r.FormValue("id")

	fmt.Println(vehicleNumber, "<- challan checking for vehicle number")
	// 

	fmt.Println(time.Since(now), "<- time took for challan API")
}

func VehicleInformation(w http.ResponseWriter, r *http.Request){
	now := time.Now()
	vehicleNumber := r.FormValue("id")

	fmt.Println(vehicleNumber, "<- checking vehicle information for vehicle number")
	// 

	fmt.Println(time.Since(now), "<- time took for challan API")
}