docker build --force-rm --tag wex:0.1.0 .
docker run --name wex --network wex-network -p 3060:3060 wex:0.1.0