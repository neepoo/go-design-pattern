package builder

import (
    "testing"

   	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T){
    manufacturingComplex := ManufacturingDirector{}

    carBuilder := new(CarBuilder)
    manufacturingComplex.SetBuilder(carBuilder)
    manufacturingComplex.Construct()
    car := carBuilder.GetVehicle()
    t.Logf("%#v\n", car)
    require.Equal(t, 4, car.Wheels)
    require.Equal(t, 5, car.Seats)
    require.Equal(t, "Car", car.Structure)

    bikeBuilder := new(BikeBuilder)
    manufacturingComplex.SetBuilder(bikeBuilder)
    manufacturingComplex.Construct()
    bike := bikeBuilder.GetVehicle()

    require.Equal(t, 2, bike.Wheels)
    require.Equal(t, 2, bike.Seats)
    require.Equal(t, "Motorbike", bike.Structure)
}