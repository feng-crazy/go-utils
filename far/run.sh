#!bin/bash

set -ex

#./far find -f ../../single_export~aiztggzj-1.0.0-c11_delivery~1667184940604/config/project_app.json -e "iregistry.baidu-int.com/\S+\/\S+" -o output/far.out.json

#./far find -d ../../single_export~aiztggzj-1.0.0-c11_delivery~1667184940604 -e "iregistry.baidu-int.com/\S+\/\S+" -o output/far.out.json

./far find -d ./ -e "iregistry.baidu-int.com/\S+\/\S+" -o output/far.out.json

./far replace  -i output/far.out.json

./far replace  -i output/far.out.json