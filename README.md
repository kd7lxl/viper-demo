# Usage

To reproduce https://github.com/spf13/viper/issues/608, run `make`.

```
$ make
go build -o viper-demo .
./viper-demo --tags=location=Tacoma
plag: &map[location:Tacoma]
viper: map[]

./viper-demo --tags=location=Tacoma,size=medium
plag: &map[location:Tacoma size:medium]
viper: map[]

./viper-demo --tags=location=Tacoma --tags=county=US
plag: &map[county:US location:Tacoma]
viper: map[]

./viper-demo --tags=location=Tacoma,size=medium --tags=county=US
plag: &map[county:US location:Tacoma size:medium]
viper: map[]

```
Notice that Viper loads an empty map.
